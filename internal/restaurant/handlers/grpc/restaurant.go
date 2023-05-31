package handlers

import (
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/contracts-v0.3.0/pkg/contracts/restaurant"
)

type RestaurantService struct {
	pb.UnimplementedProductServiceServer
	pb.UnimplementedMenuServiceServer
	pb.UnimplementedOrderServiceServer

	repoR RestaurantRepository
}

func NewRestaurantService(rp RestaurantRepository) *RestaurantService {
	return &RestaurantService{repoR: rp}
}

type RestaurantRepository interface {
	CreateProduct(*pb.Product) error
	GetProductList() (*pb.GetProductListResponse, error)
	CreateMenu() (*pb.Menu, error)
	GetMenu(time.Time) *pb.Menu
	GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice)
}
