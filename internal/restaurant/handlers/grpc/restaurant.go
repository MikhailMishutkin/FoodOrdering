package handlers

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"time"
)

type RestaurantService struct {
	restaurant.UnimplementedProductServiceServer
	restaurant.UnimplementedMenuServiceServer
	restaurant.UnimplementedOrderServiceServer

	rSer RestaurantServicer
}

func NewRestaurantService(rs RestaurantServicer) *RestaurantService {
	return &RestaurantService{
		rSer: rs,
	}
}

type RestaurantServicer interface {
	CreateProduct(product *types.Product) error
	GetProductList() ([]*types.Product, error)
	CreateMenu(create *types.MenuCreate) error
	GetMenu(time.Time) (*types.Menu, error)
	GetOrderList() ([]*types.OrderItem, []*types.OrderByOffice, error)
}
