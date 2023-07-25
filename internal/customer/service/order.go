package service

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"log"
)

func (cu *CustomerUsecase) CreateOrder(request *types.OrderRequest) error {
	log.Println("CreateOrder service was invoked")

	err := cu.repoC.CreateOrder(request)

	return err
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
