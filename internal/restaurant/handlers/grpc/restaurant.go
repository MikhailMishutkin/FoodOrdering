package handlers

import (
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

type RestaurantService struct {
	pb.UnimplementedProductServiceServer
	pb.UnimplementedMenuServiceServer
	pb.UnimplementedOrderServiceServer

	repoR RestaurantRepository
}

func NewRestaurantService(rp RestaurantRepository) *RestaurantService {
	//	log.Println("check repoR", rp)
	return &RestaurantService{repoR: rp}
}

type RestaurantRepository interface {
	CreateProduct(*pb.Product) error
	GetProductList() (*pb.GetProductListResponse, error)
	CreateMenu() (*pb.Menu, error)
	GetMenu(time.Time) (*pb.Menu, error)
	GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice)
}
