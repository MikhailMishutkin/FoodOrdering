package service

import (
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"log"
)

// GetActualMenu from restaurant DB
func (cu *CustomerUsecase) GetActualMenu(res *restaurant.GetMenuResponse) (amr *pb.GetActualMenuResponse, err error) {

	amr = &pb.GetActualMenuResponse{
		Salads:    ProductConv(res.Menu.Salads),
		Garnishes: ProductConv(res.Menu.Garnishes),
		Meats:     ProductConv(res.Menu.Meats),
		Soups:     ProductConv(res.Menu.Soups),
		Drinks:    ProductConv(res.Menu.Drinks),
		Desserts:  ProductConv(res.Menu.Desserts),
	}

	return amr, nil
}

func (cu *CustomerUsecase) CreateOrder(in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Println("CreateOrder service was invoked")

	err := cu.repoC.CreateOrder(in)
	return &pb.CreateOrderResponse{}, err
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

// m := dataMap
// salads := menu.Salads
// for _, v := range salads {
// 	m[v.Uuid] = v
// }

//go natsSubscriber()
