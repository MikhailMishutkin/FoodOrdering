package repository

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

// TODO: ошибки!!!!
func (r *RestaurantRepo) GetOrderList(
	date time.Time,
	offices []*types.Office,
	users []*types.User,
) (
	[]*types.OrderItem,
	[]*types.OrderByOffice,
	error,
) {

	log.Println("GetOrderList repository was invoked")
	//slice of OrderItem
	q := `SELECT  O.product_id, (SELECT name FROM product WHERE uuid = O.product_id), SUM(O.count) AS total
FROM orders AS O
WHERE O.on_date = $1
GROUP BY O.product_id`

	rows, err := r.DB.Query(q, date)
	if err != nil {
		return nil, nil, fmt.Errorf("Error while select items for slice of OrderItem (GetOrderList): %v\n", err) //TODO
	}

	slOI := []*types.OrderItem{}
	for rows.Next() {
		order := &types.OrderItem{}
		if err := rows.Scan(
			&order.ProductUuid,
			&order.ProductName,
			&order.Count,
		); err != nil {
			fmt.Errorf("Something wrong with rows.Scan while scanning OrderItem for TotalOrders in GetOrderList: %s", err)
		}
		slOI = append(slOI, order)
	}
	// write offices to db
	for _, v := range offices {
		q = `INSERT INTO office  (id, office_name, office_address) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`
		_, err = r.DB.Exec(q,
			v.Uuid,
			v.Name,
			v.Address,
		)
		if err != nil {
			fmt.Errorf("Can't insert into db offices data (GetOrderList): %v\n", err)

		}
	}
	//write users to db
	for _, v := range users {
		q = `INSERT INTO users (id, user_name, office_uuid) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`
		_, err = r.DB.Exec(q,
			v.Uuid,
			v.Name,
			v.OfficeUuid,
		)
		if err != nil {
			fmt.Errorf("Can't insert into db users data (GetOrderList): %v\n", err)

		}
	}

	//slice of OrdersByCompany
	rows1, err := r.DB.Query("SELECT * FROM office")
	if err != nil {
		fmt.Errorf("can't get offices from db in GetOrderList repository: %v\n", err)
	}
	slOrderByOffice := []*types.OrderByOffice{}
	for rows1.Next() {
		office := &types.OrderByOffice{}
		if err = rows1.Scan(&office.OfficeUuid, &office.OfficeName, &office.OfficeAddress); err != nil {
			fmt.Errorf(
				`trouble with rows1.Scan (GetOrderList Restaurant repo,
						for slice of OrdersByCompany): %v`,
				err,
			)
		}
		q1 := `SELECT O.product_id, (SELECT name FROM product WHERE uuid = O.product_id), SUM(O.count) AS total
				FROM orders AS O
				WHERE O.on_date = $1 and O.user_uuid IN (SELECT id
                                                  FROM users
                                                  WHERE office_uuid = $2)
				GROUP BY O.product_id`
		rows2, err := r.DB.Query(q1, date, office.OfficeUuid)
		if err != nil {
			return nil, nil, fmt.Errorf("Error with select from db summ of product by office (GetOrderList): %v\n", err)
		}
		slOrderItemByOffice := []*types.OrderItem{}
		for rows2.Next() {
			oI := &types.OrderItem{}
			if err = rows2.Scan(&oI.ProductUuid, &oI.ProductName, &oI.Count); err != nil {
				return nil, nil, fmt.Errorf("Something wrong with rows2.Scan while scanning OrderItem for OrdersByCompany (GetOrderList): %v\n", err)
			}
			slOrderItemByOffice = append(slOrderItemByOffice, oI)
		}
		office.Result = slOrderItemByOffice
		slOrderByOffice = append(slOrderByOffice, office)

	}

	return slOI, slOrderByOffice, err
}

func (r *RestaurantRepo) ReceiveOrder(order *types.OrderRequest) error {
	log.Println("ReceiveOrder (order_repo restaurant) was invoked")

	var slOI []*types.OrderItem

	slOI = append(slOI, order.Salads...)
	slOI = append(slOI, order.Garnishes...)
	slOI = append(slOI, order.Meats...)
	slOI = append(slOI, order.Soups...)
	slOI = append(slOI, order.Drinks...)
	slOI = append(slOI, order.Desserts...)

	for _, v := range slOI {
		res2B, _ := json.Marshal(v)
		fmt.Println(string(res2B))
		if v.Count != 0 && v.ProductUuid != 0 {
			_, err := r.DB.Exec(
				"INSERT INTO orders (user_uuid, product_id, count) VALUES ($1, $2, $3)", //ON CONFLICT DO UPDATE SET count = EXCLUDED.count
				order.UserUuid,
				v.ProductUuid,
				v.Count,
			)
			if err != nil {
				fmt.Errorf("Can't INSERT order into restaurant db: %v\n", err)
			}
		}
	}
	return nil
}
