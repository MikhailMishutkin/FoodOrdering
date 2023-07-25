package repository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

// TODO: добавить вариант если названия нет в бд
// TODO: сделать сложный ключ для UUID, возможно ещё одну таблицу инкремент-дата
func (r *RestaurantRepo) CreateMenu(mc *types.MenuCreate) error {
	m := &types.Menu{}
	var err error
	var slp []string
	slp = concantenateProducts(mc.Salads, slp)
	slp = concantenateProducts(mc.Garnishes, slp)
	slp = concantenateProducts(mc.Meats, slp)
	slp = concantenateProducts(mc.Soups, slp)
	slp = concantenateProducts(mc.Drinks, slp)
	slp = concantenateProducts(mc.Desserts, slp)
	for _, v := range slp {
		var id, t int
		err := r.DB.QueryRow("SELECT uuid, type_id  FROM product WHERE name = $1", v).Scan(&id, &t)
		if err != nil {
			fmt.Errorf("Error to get id and type from db: %s", err)
		}
		err = r.DB.QueryRow("INSERT INTO menu (on_date, opening_record_at, closing_record_at, product_id, prod_type_m) VALUES ($1, $2, $3, $4, $5) RETURNING uuid",
			mc.OnDate,
			mc.OpenAt,
			mc.ClosedAt,
			id,
			t,
		).Scan(&m.Uuid)
		if err != nil {
			return fmt.Errorf("Can't INSERT into db menu data: %v\n", err)
		}
	}

	return err
}

func (r *RestaurantRepo) GetMenu(t time.Time) (*types.Menu, error) {

	log.Printf("GetMenu Repository was invoked")
	menu := &types.Menu{}
	var slp []*types.Product
	rows, err := r.DB.Query("SELECT product_id FROM menu WHERE on_date = $1", t)
	if err != nil {
		return nil, fmt.Errorf("Error to get id and type from db: %s", err)
	}

	for rows.Next() {
		p := &types.Product{}
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("trouble with rows.Next: %s", err)
		}
		err = r.DB.QueryRow("SELECT uuid, name, description, type_id, weight, price, created_at FROM product WHERE uuid = $1",
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

	err = r.DB.QueryRow("SELECT opening_record_at, closing_record_at, created_at FROM menu WHERE on_date = $1", t).Scan(
		&menu.OpenAt,
		&menu.ClosedAt,
		&menu.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("Error to get time from db: %s", err)
	}

	return menu, nil
}

func concantenateProducts(sl []string, slp []string) []string {
	for _, v := range sl {
		slp = append(slp, v)
	}
	return slp
}
