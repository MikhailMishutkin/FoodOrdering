package handlers_customer

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
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
	CreateOffice(*types.Office) error
	GetOfficeList() ([]*types.Office, error)
	CreateUser(*types.User) error
	GetUserList(int) ([]*types.User, error)
	//GetActualMenu(menu *types.Menu) (, error)
	CreateOrder(request *types.OrderRequest) error
}
