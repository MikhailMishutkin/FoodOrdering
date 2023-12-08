package service

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
)

type CustomerUsecase struct {
	repoC CustomerRepositorier
	nr    NatsCustomerRepositorier
}

func NewCustomerUsecase(r CustomerRepositorier, nr NatsCustomerRepositorier) *CustomerUsecase {
	return &CustomerUsecase{
		repoC: r,
		nr:    nr,
	}
}

type CustomerRepositorier interface {
	CreateOffice(office *types.Office) error
	GetOfficeList() ([]*types.Office, error)
	CreateUser(user *types.User) error
	GetUserList(int) ([]*types.User, error)
}

type NatsCustomerRepositorier interface {
	NatsPublisher(order *types.OrderRequest) error
}
