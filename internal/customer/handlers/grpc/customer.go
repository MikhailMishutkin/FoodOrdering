package handlers_customer

import (
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

type CustomerService struct {
	pb.UnimplementedOfficeServiceServer
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedUserServiceServer

	client restaurant.MenuServiceClient
	cs     CustomerServicer
}

//func NewCustomerService(cs CustomerServicer) *CustomerService {
//	return &CustomerService{cs: cs}
//}

func New(client restaurant.MenuServiceClient, cs CustomerServicer) *CustomerService {
	//fmt.Println("функция нью пакета хандлерс_кастомер: ", &client)
	return &CustomerService{
		cs:     cs,
		client: client,
	}
}

type CustomerServicer interface {
	CreateOffice(*pb.Office) error
	GetOfficeList() ([]*pb.Office, error)
	CreateUser(*pb.User) error
	GetUserList(*pb.GetUserListRequest) (*pb.GetUserListResponse, error)
	GetActualMenu(*restaurant.GetMenuResponse) (*pb.GetActualMenuResponse, error)
	CreateOrder(*pb.CreateOrderRequest) (*pb.CreateOrderResponse, error)
}
