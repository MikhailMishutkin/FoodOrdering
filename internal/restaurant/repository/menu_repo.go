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

//func typeSelect(ti int) string {
//	switch ti {
//	case 1:
//		return "INSERT INTO menu (product_type_m) VALUES ($1) WHERE uuid= id "
//	case 2:
//		return "INSERT INTO menu (garnishes) VALUES ($1) RETURNING uuid"
//	case 3:
//		return "INSERT INTO menu (meats) VALUES ($1) RETURNING uuid"
//	case 4:
//		return "INSERT INTO menu (soups) VALUES ($1) RETURNING uuid"
//	case 5:
//		return "INSERT INTO menu (drinks) VALUES ($1) RETURNING uuid"
//	case 6:
//		return "INSERT INTO menu (desserts) VALUES ($1) RETURNING uuid"
//
//	}
//}

//func extractProductByName(sl []string) {
//	ndb, _ := NewDB()
//	rr := NewRestaurantRepo(ndb)
//	for _, v := range sl {
//		rows, err := rr.db.Query("SELECT uuid FROM product WHERE name = $1", v)
//		if err != nil {
//			return fmt.Errorf("Error to get name of product from db: %v\n", err)
//		}
//		defer rows.Close()
//		for rows.Next() {
//			if err := rows.Scan(&uuid); err != nil {
//				r.logger.Printf("", err)
//				return nil, err
//			}
//	}
//
//}

//create menu without db
//m := make(map[string]*pb.Product) //настройка бд
//
//var Salads, Garnishes, Meats, Soups, Drinks, Desserts []*pb.Product
//for _, v := range m {
//	switch {
//	case v.Type == pb.ProductType_PRODUCT_TYPE_SALAD:
//		Salads = append(Salads, v)
//	case v.Type == pb.ProductType_PRODUCT_TYPE_GARNISH:
//		Garnishes = append(Garnishes, v)
//	case v.Type == pb.ProductType_PRODUCT_TYPE_MEAT:
//		Meats = append(Meats, v)
//	case v.Type == pb.ProductType_PRODUCT_TYPE_SOUP:
//		Soups = append(Soups, v)
//	case v.Type == pb.ProductType_PRODUCT_TYPE_DRINK:
//		Drinks = append(Drinks, v)
//	default:
//		Desserts = append(Desserts, v)
//	}
//}
//t := DateConv(time.Now())
//
//t1 := t.AddDate(0, 0, 1)
//ts := timestamppb.New(t1)
//
//t2 := t1.Add(11 * time.Hour)
//ts1 := timestamppb.New(t2)
//t3 := t1.Add(21 * time.Hour)
//ts2 := timestamppb.New(t3)
//
//menu := &pb.Menu{
//Uuid:            RandomID(),
//OnDate:          ts,
//OpeningRecordAt: ts1,
//ClosingRecordAt: ts2,
//Salads:          Salads,
//Garnishes:       Garnishes,
//Meats:           Meats,
//Soups:           Soups,
//Drinks:          Drinks,
//Desserts:        Desserts,
//CreatedAt:       timestamppb.Now(),
//}
// getMenu without db
//data, err := os.OpenFile("menu.json", os.O_RDONLY, 0644)
//if err != nil {
//log.Fatal("can't read menu.json", err)
//
//}
//defer data.Close()
//
//m, err := io.ReadAll(data)
//if err != nil {
//log.Println("Can't read data from menu.json: ", err)
//return &menuInst, err
//}
//menu := &pb.Menu{}
//
//err = json.Unmarshal(m, menu)
//if err != nil {
//log.Fatal("cannot unmarshall menu", err)
//}
//fmt.Println("поле структуры меню OnDate: ", d)
//t1 := DateConv(t)
//if d == t1 {
//
//return menu, nil
//}
