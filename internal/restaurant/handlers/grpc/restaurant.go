package handlers

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
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
	CreateProduct(product *types.Product) error
	GetProductList() ([]*types.Product, error)
	CreateMenu(create *types.MenuCreate) error
	GetMenu(time.Time) (*types.Menu, error)
	GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice)
}
