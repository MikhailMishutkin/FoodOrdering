package service

import rest "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"

type CustomerUsecase struct {
	rest.UnimplementedOrderServiceServer
	client rest.MenuServiceClient
	repoC  CustomerRepository
}

func NewCustomerUsecase(r CustomerRepository) *CustomerUsecase {
	return &CustomerUsecase{repoC: r}
}

func New(client rest.MenuServiceClient) *CustomerUsecase {
	return &CustomerUsecase{
		client: client,
	}
}

type CustomerRepository interface {
	CreateOrder()
	CreateOffice()
}
