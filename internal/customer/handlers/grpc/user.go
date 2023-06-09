package handlers_customer

import (
	"context"
	"fmt"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
)

//TODO: Uuid для пользователя уже из базы, проверка на существование в случае ручного создания
func (s *CustomerService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := new(pb.User)
	//if in.Name == "" {
	//	user = gen.NewUser()
	//}

	fmt.Println(in.Name, user)
	user.Name = in.Name
	user.Uuid = in.OfficeUuid
	err := s.cs.CreateUser(user)

	return &pb.CreateUserResponse{}, err
}

func (s *CustomerService) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {

	res, err := s.cs.GetUserList(in)

	return res, err
}
