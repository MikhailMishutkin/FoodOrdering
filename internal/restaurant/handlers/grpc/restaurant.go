package handlers

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

type RestaurantService struct {
	pb.UnimplementedProductServiceServer
	pb.UnimplementedMenuServiceServer
	pb.UnimplementedOrderServiceServer

	rSer      RestaurantServicer
	OffClient customer.OfficeServiceClient
	UsClient  customer.UserServiceClient
}

func NewRestaurantService(
	rs RestaurantServicer,
	OffClient customer.OfficeServiceClient,
	UsClient customer.UserServiceClient,
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
	//GetOrder(order *types.OrderRequest) error
}
