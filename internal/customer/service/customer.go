package service

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	rest "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

type CustomerUsecase struct {
	rest.UnimplementedOrderServiceServer
	client rest.MenuServiceClient
	repoC  CustomerRepositorier
}

func NewCustomerUsecase(r CustomerRepositorier) *CustomerUsecase {
	return &CustomerUsecase{repoC: r}
}

func New(client rest.MenuServiceClient) *CustomerUsecase {
	return &CustomerUsecase{
		client: client,
	}
}

type CustomerRepositorier interface {
	CreateOffice(office *types.Office) error
	GetOfficeList() ([]*types.Office, error)
	CreateUser(user *types.User) error
	GetUserList(int) ([]*types.User, error)
	CreateOrder(request *types.OrderRequest) error
}
