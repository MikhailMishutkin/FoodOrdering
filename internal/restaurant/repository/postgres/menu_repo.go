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
		"SELECT uuid, type_id  FROM product WHERE name = $1 and created_at = $2",
		name,
		date,
	).Scan(&id, &t)
	log.Println(id, t)
	if err != nil {
		fmt.Errorf("Error to get id and type from db: %s", err)
	}
	return id, t, nil
}

// ...
func (r *RestaurantRepo) CreateDate(date time.Time) (int, error) {
	var id int
	err := r.DB.QueryRow(
		context.Background(),
		"INSERT INTO menu_date (date) VALUES ($1) RETURNING uuid",
		date,
	).Scan(&id)

	return id, err
}

// TODO: добавить вариант если названия нет в бд
// TODO: сделать сложный ключ для UUID, возможно ещё одну таблицу инкремент-дата
func (r *RestaurantRepo) CreateMenu(id int, menuId int, pt int) error {
	m := &types.Menu{}

	err := r.DB.QueryRow(context.Background(),
		"INSERT INTO menu (menu_id, product_id, prod_type_m) VALUES ($1, $2, $3) RETURNING uuid",
		id,
		menuId,
		pt,
	).Scan(&m.Uuid)
	if err != nil {
		return fmt.Errorf("Can't INSERT into db menu data: %v\n", err)
	}

	return err
}

func (r *RestaurantRepo) GetMenu(t time.Time) (*types.Menu, error) {

	log.Printf("GetMenu Repository was invoked %v\n", t)

	menu := &types.Menu{}
	var slp []*types.Product
	rows, err := r.DB.Query(context.Background(), "SELECT product_id FROM menu WHERE on_date = $1", t)
	if err != nil {
		return nil, fmt.Errorf("Error to get id and type from db: %s", err)
	}

	for rows.Next() {
		p := &types.Product{}
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("trouble with rows.Next: %s", err)
		}
		err = r.DB.QueryRow(context.Background(), "SELECT uuid, name, description, type_id, weight, price, created_at FROM product WHERE uuid = $1",
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
			fmt.Errorf("Error to get id and type from db: %s", err)
		}
		slp = append(slp, p)
	}
	for _, v := range slp {
		switch {
		case v.Type == 1:
			menu.Salads = append(menu.Salads, v)
		case v.Type == 2:
			menu.Garnishes = append(menu.Garnishes, v)
		case v.Type == 3:
			menu.Meats = append(menu.Meats, v)
		case v.Type == 4:
			menu.Soups = append(menu.Soups, v)
		case v.Type == 5:
			menu.Drinks = append(menu.Drinks, v)
		default:
			menu.Desserts = append(menu.Desserts, v)
		}
	}

	err = r.DB.QueryRow(context.Background(), "SELECT opening_record_at, closing_record_at, created_at FROM menu WHERE on_date = $1", t).Scan(
		&menu.OpenAt,
		&menu.ClosedAt,
		&menu.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("Error to get time from db: %s", err)
	}

	return menu, nil
}
