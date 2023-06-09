package repository

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (r *RestaurantRepo) CreateMenu() (*pb.Menu, error) {
	m := dataMap

	var Salads, Garnishes, Meats, Soups, Drinks, Desserts []*pb.Product
	for _, v := range m {
		switch {
		case v.Type == pb.ProductType_PRODUCT_TYPE_SALAD:
			Salads = append(Salads, v)
		case v.Type == pb.ProductType_PRODUCT_TYPE_GARNISH:
			Garnishes = append(Garnishes, v)
		case v.Type == pb.ProductType_PRODUCT_TYPE_MEAT:
			Meats = append(Meats, v)
		case v.Type == pb.ProductType_PRODUCT_TYPE_SOUP:
			Soups = append(Soups, v)
		case v.Type == pb.ProductType_PRODUCT_TYPE_DRINK:
			Drinks = append(Drinks, v)
		default:
			Desserts = append(Desserts, v)
		}
	}

	t := DateConv(time.Now())

	t1 := t.AddDate(0, 0, 1)
	ts := timestamppb.New(t1)

	t2 := t1.Add(11 * time.Hour)
	ts1 := timestamppb.New(t2)
	t3 := t1.Add(21 * time.Hour)
	ts2 := timestamppb.New(t3)

	menu := &pb.Menu{
		Uuid:            RandomID(),
		OnDate:          ts,
		OpeningRecordAt: ts1,
		ClosingRecordAt: ts2,
		Salads:          Salads,
		Garnishes:       Garnishes,
		Meats:           Meats,
		Soups:           Soups,
		Drinks:          Drinks,
		Desserts:        Desserts,
		CreatedAt:       timestamppb.Now(),
	}

	return menu, nil
}

func (r *RestaurantRepo) GetMenu(t time.Time) (*pb.Menu, error) {
	log.Printf("GetMenu Repository was invoked")
	menuInst := pb.Menu{}
	data, err := os.OpenFile("menu.json", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("can't read menu.json", err)

	}
	defer data.Close()

	m, err := io.ReadAll(data)
	if err != nil {
		log.Println("Can't read data from menu.json: ", err)
		return &menuInst, err
	}
	menu := &pb.Menu{}

	err = json.Unmarshal(m, menu)
	if err != nil {
		log.Fatal("cannot unmarshall menu", err)
	}

	d := menu.OnDate.AsTime()
	//fmt.Println("поле структуру меню OnDate: ", d)
	t1 := DateConv(t)
	if d == t1 {

		return menu, nil
	}

	return &menuInst, nil
}
