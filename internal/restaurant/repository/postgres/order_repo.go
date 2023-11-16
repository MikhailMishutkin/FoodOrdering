package repository

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

// TODO: ошибки!!!!
func (r *RestaurantRepo) GetTotalOrders(date time.Time) ([]*types.OrderItem, error) {
	log.Println("GetTotalOrders was invoked ")
	q := `SELECT  O.product_id, (SELECT name FROM product WHERE uuid = O.product_id), SUM(O.count) AS total
FROM orders AS O
WHERE O.on_date = $1
GROUP BY O.product_id`

	rows, err := r.DB.Query(context.Background(), q, date)
	if err != nil {
		return nil, fmt.Errorf("Error while select items for slice of OrderItem (GetOrderList): %v\n", err) //TODO
	}
	var slOI []*types.OrderItem
	for rows.Next() {
		order := &types.OrderItem{}
		if err := rows.Scan(
			&order.ProductUuid,
			&order.ProductName,
			&order.Count,
		); err != nil {
			return nil, fmt.Errorf("Something wrong with rows.Scan while scanning OrderItem for TotalOrders in GetOrderList: %s", err)
		}

		slOI = append(slOI, order)
	}
	return slOI, nil
}

// ....
func (r *RestaurantRepo) GetOfficesList() ([]*types.OrderByOffice, error) {
	var slOrderByOffice []*types.OrderByOffice
	rows1, err := r.DB.Query(context.Background(), "SELECT * FROM office")
	if err != nil {
		return nil, fmt.Errorf("can't get offices from db in GetOrderList repository: %v\n", err)
	}

	for rows1.Next() {
		office := &types.OrderByOffice{}
		if err = rows1.Scan(&office.OfficeUuid, &office.OfficeName, &office.OfficeAddress); err != nil {
			fmt.Errorf(
				`trouble with rows1.Scan (GetOfficesList): %v`,
				err,
			)
		}
		slOrderByOffice = append(slOrderByOffice, office)
	}
	return slOrderByOffice, err
}

// ...
func (r *RestaurantRepo) GetOrdersByOffice(date time.Time, id int) ([]*types.OrderItem, error) {

	log.Println("GetOrderList repository was invoked")
	q1 := `SELECT O.product_id, (SELECT name FROM product WHERE uuid = O.product_id), SUM(O.count) AS total
				FROM orders AS O
				WHERE O.on_date = $1 and O.user_uuid IN (SELECT id
	                                             FROM users
	                                             WHERE office_uuid = $2)
				GROUP BY O.product_id`
	rows, err := r.DB.Query(context.Background(), q1, date, id)
	if err != nil {
		return nil, fmt.Errorf("Error with select from db summ of product by office (GetOrderList): %v\n", err)
	}

	var slOrderItemByOffice []*types.OrderItem
	for rows.Next() {
		oI := &types.OrderItem{}
		if err = rows.Scan(&oI.ProductUuid, &oI.ProductName, &oI.Count); err != nil {
			return nil, fmt.Errorf("Something wrong with rows2.Scan while scanning OrderItem for OrdersByCompany (GetOrderList): %v\n", err)
		}
		slOrderItemByOffice = append(slOrderItemByOffice, oI)
	}

	return slOrderItemByOffice, nil
}

// ...
func (r *RestaurantRepo) ReceiveOrder(slOI []*types.OrderItem, userUuid int) error {
	log.Println("ReceiveOrder (order_repo restaurant) was invoked")

	for _, v := range slOI {
		if v.Count != 0 && v.ProductUuid != 0 {
			_, err := r.DB.Exec(context.Background(),
				"INSERT INTO orders (user_uuid, product_id, count) VALUES ($1, $2, $3)", //ON CONFLICT DO UPDATE SET count = EXCLUDED.count
				userUuid,
				v.ProductUuid,
				v.Count,
			)
			if err != nil {
				log.Printf("Can't INSERT order into restaurant db: %v\n", err)
			}
		}
	}
	return nil
}

// ...
func (r *RestaurantRepo) SaveOfficeList(offices []*types.Office) error {
	// write offices to db
	for _, v := range offices {
		q := `INSERT INTO office  (id, office_name, office_address) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`
		_, err := r.DB.Exec(context.Background(),
			q,
			v.Uuid,
			v.Name,
			v.Address,
		)
		if err != nil {
			return fmt.Errorf("Can't insert into db offices data (SaveOfficeList): %v\n", err)

		}
	}
	return nil
}

// ...
func (r *RestaurantRepo) SaveUserList(users []*types.User) error {
	//write users to db
	for _, v := range users {
		q := `INSERT INTO users (id, user_name, office_uuid) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING`
		_, err := r.DB.Exec(context.Background(),
			q,
			v.Uuid,
			v.Name,
			v.OfficeUuid,
		)
		if err != nil {
			fmt.Errorf("Can't insert into db users data (SaveUserList): %v\n", err)

		}
	}
	return nil
}
