package handlers_customer

import (
	"context"
	"log"

	cl "github.com/MikhailMishutkin/FoodOrdering/cmd/customer/customer_app"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

func (s *CustomerService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return nil, nil
}

func (s *CustomerService) GetActualMenu(ctx context.Context, in *pb.GetActualMenuRequest) (*pb.GetActualMenuResponse, error) {
	log.Println("GetActualMenu was invoked")
	rmr, err := cl.Сonn()
	if err != nil {
		log.Println("client.Conn error", err)
		return nil, err
	}

	amr := &pb.GetActualMenuResponse{
		Salads:    ProductConv(rmr.Menu.Salads),
		Garnishes: ProductConv(rmr.Menu.Garnishes),
		Meats:     ProductConv(rmr.Menu.Meats),
		Soups:     ProductConv(rmr.Menu.Soups),
		Drinks:    ProductConv(rmr.Menu.Drinks),
		Desserts:  ProductConv(rmr.Menu.Desserts),
	}
	return amr, nil
}

func ProductConv(p []*restaurant.Product) []*pb.Product {
	var sl []*pb.Product
	for _, v := range p {
		a := &pb.Product{
			Uuid:        v.Uuid,
			Name:        v.Name,
			Description: v.Description,
			Type:        pb.CustomerProductType(v.Type),
			Weight:      v.Weight,
			Price:       v.Price,
			CreatedAt:   v.CreatedAt,
		}
		sl = append(sl, a)
	}
	return sl
}
