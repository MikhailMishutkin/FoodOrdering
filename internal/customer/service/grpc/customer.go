package service

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
)

type CustomerUsecase struct {
	restaurant.UnimplementedOrderServiceServer
	client restaurant.MenuServiceClient
	repoC  CustomerRepositorier
}

func NewCustomerUsecase(r CustomerRepositorier) *CustomerUsecase {
	return &CustomerUsecase{repoC: r}
}

func New(client restaurant.MenuServiceClient) *CustomerUsecase {
	return &CustomerUsecase{
		client: client,
	}
}

type CustomerRepositorier interface {
	CreateOffice(office *types.Office) error
	GetOfficeList() ([]*types.Office, error)
	CreateUser(user *types.User) error
	GetUserList(int) ([]*types.User, error)
}
