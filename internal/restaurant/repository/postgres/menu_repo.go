package repository

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"time"
)

// ...
func (r *RestaurantRepo) SelectProductByName(name string, date time.Time) (int, int, error) {
	var id, t int
	err := r.DB.QueryRow(context.Background(),
		"SELECT uuid, type_id  FROM product WHERE name = $1 AND created_at::date = $2::date",
		name,
		date,
	).Scan(&id, &t)
	log.Println(id, t)
	if err != nil {
		return 0, 0, fmt.Errorf("Error to get id and type from db: %s", err)
	}
	return id, t, nil
}

// ...
func (r *RestaurantRepo) CreateMenuDate(mc *types.MenuCreate) (int, error) {
	var id int
	err := r.DB.QueryRow(
		context.Background(),
		"INSERT INTO menu_date (date, opening_record_at, closing_record_at) VALUES ($1, $2, $3) RETURNING uuid",
		mc.OnDate,
		mc.OpenAt,
		mc.ClosedAt,
	).Scan(&id)

	return id, err
}

// TODO: добавить вариант если названия нет в бд
// TODO: сделать сложный ключ для UUID, возможно ещё одну таблицу инкремент-дата
func (r *RestaurantRepo) CreateMenu(menuId int, id int, pt int) error {
	//m := &types.Menu{}

	_, err := r.DB.Exec(context.Background(),
		"INSERT INTO menu (menu_id, product_id, prod_type_m) VALUES ($1, $2, $3) RETURNING uuid",
		menuId,
		id,
		pt,
	)

	if err != nil {
		return fmt.Errorf("Can't INSERT into db menu data: %v\n", err)
	}

	return err
}

// ...
func (r *RestaurantRepo) GetMenuId(t time.Time) (int, error) {
	var id int
	err := r.DB.QueryRow(context.Background(),
		"SELECT uuid FROM menu_date WHERE date::date = $1",
		t,
	).Scan(&id)

	return id, err
}

// ...
func (r *RestaurantRepo) GetProductId(id int) (productsId []int, err error) {

	rows, err := r.DB.Query(
		context.Background(),
		"SELECT product_id FROM menu WHERE menu_id = $1",
		id,
	)
	if err != nil {
		return nil, fmt.Errorf("Error to get product_id from db: %s", err)
	}

	for rows.Next() {
		var pid int
		if err := rows.Scan(&pid); err != nil {
			return nil, fmt.Errorf("trouble with rows.Next: %s", err)
		}
		productsId = append(productsId, pid)
	}
	return productsId, err
}

// ...
func (r *RestaurantRepo) GetMenu(id int) (*types.Product, error) {

	log.Println("GetMenu Repository was invoked")

	p := &types.Product{}
	err := r.DB.QueryRow(context.Background(),
		"SELECT uuid, name, description, type_id, weight, price, created_at FROM product WHERE uuid = $1",
		id,
	).Scan(
		&p.Uuid,
		&p.Name,
		&p.Descript,
		&p.Type,
		&p.Weight,
		&p.Price,
		&p.CreatedAt,
	)
	if err != nil {
		fmt.Errorf("Error to get product fields from db: %s", err)
	}

	return p, err
}

// ...
func (r *RestaurantRepo) GetTimes(id int) (tO time.Time, tC time.Time, tCr time.Time, err error) {
	err = r.DB.QueryRow(
		context.Background(),
		"SELECT opening_record_at, closing_record_at, created_at FROM menu_date WHERE uuid = $1", id).Scan(
		&tO,
		&tC,
		&tCr,
	)
	if err != nil {
		return tO, tC, tCr, fmt.Errorf("Error to get time from db: %s", err)
	}
	return tO, tC, tCr, err
}
