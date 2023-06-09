package handlers

import (
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

type RestaurantService struct {
	pb.UnimplementedProductServiceServer
	pb.UnimplementedMenuServiceServer
	pb.UnimplementedOrderServiceServer

	rSer RestaurantServicer
}

func NewRestaurantService(rs RestaurantServicer) *RestaurantService {
	return &RestaurantService{rSer: rs}
}

type RestaurantServicer interface {
	CreateProduct(*pb.Product) error
	GetProductList() (*pb.GetProductListResponse, error)
	CreateMenu() (*pb.Menu, error)
	GetMenu(time.Time) (*pb.Menu, error)
	GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice)
}
