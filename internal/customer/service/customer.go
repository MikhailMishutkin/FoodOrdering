package service

import (
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	rest "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

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
	CreateOffice(office *pb.Office) error
	GetOfficeList() ([]*pb.Office, error)
	CreateUser(user *pb.User) error
	GetUserList(string) (*pb.GetUserListResponse, error)
	CreateOrder(*pb.CreateOrderRequest) error
}
