package handlers_customer

import (
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/contracts-v0.3.0/pkg/contracts/customer"
)

type CustomerService struct {
	pb.UnimplementedOfficeServiceServer
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedUserServiceServer

	repoC CustomerRepository
}

func NewCustomerService(rp CustomerRepository) *CustomerService {
	return &CustomerService{repoC: rp}
}

type CustomerRepository interface {
	CreateOrder()
	GetActualMenu(*pb.GetActualMenuResponse)
}
