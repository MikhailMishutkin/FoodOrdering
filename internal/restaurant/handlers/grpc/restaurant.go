package handlers

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	customer2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"time"
)

type RestaurantService struct {
	restaurant.UnimplementedProductServiceServer
	restaurant.UnimplementedMenuServiceServer
	restaurant.UnimplementedOrderServiceServer

	rSer      RestaurantServicer
	OffClient customer2.OfficeServiceClient
	UsClient  customer2.UserServiceClient
}

func NewRestaurantService(
	rs RestaurantServicer,
	OffClient customer2.OfficeServiceClient,
	UsClient customer2.UserServiceClient,
) *RestaurantService {
	return &RestaurantService{
		rSer:      rs,
		OffClient: OffClient,
		UsClient:  UsClient,
	}
}

type RestaurantServicer interface {
	CreateProduct(product *types.Product) error
	GetProductList() ([]*types.Product, error)
	CreateMenu(create *types.MenuCreate) error
	GetMenu(time.Time) (*types.Menu, error)
	GetOrderList([]*types.Office, []*types.User) ([]*types.OrderItem, []*types.OrderByOffice, error)
}
