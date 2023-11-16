package handlerscustomer

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
)

type CustomerService struct {
	customer.UnimplementedOfficeServiceServer
	customer.UnimplementedUserServiceServer
	customer.UnimplementedOrderServiceServer
	client restaurant.MenuServiceClient
	cs     CustomerServicer
}

func New(client restaurant.MenuServiceClient, cs CustomerServicer) *CustomerService {
	return &CustomerService{
		cs:     cs,
		client: client,
	}
}

type CustomerServicer interface {
	CreateOffice(*types.Office) error
	GetOfficeList() ([]*types.Office, error)
	CreateUser(*types.User) error
	GetUserList(int) ([]*types.User, error)
	CreateOrder(request *types.OrderRequest) error
}
