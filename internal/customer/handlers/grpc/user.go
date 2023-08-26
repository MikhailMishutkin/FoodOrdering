package handlerscustomer

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"log"
	"strconv"
)

// TODO: Проверка на существование в случае ручного создания
func (s *CustomerService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Println("CreateUser was invoked")
	ofId, err := strconv.Atoi(in.OfficeUuid)
	if err != nil {
		return nil, fmt.Errorf("Error in office uuid, can't convert to int: %v\n", err)
	}
	user := &types.User{
		Name:       in.Name,
		OfficeUuid: ofId,
	}

	err = s.cs.CreateUser(user)

	return &pb.CreateUserResponse{}, err
}

func (s *CustomerService) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	id, err := strconv.Atoi(in.OfficeUuid)
	if err != nil {
		return nil, fmt.Errorf("Error in office uuid(GetList), can't convert to int: %v\n", err)
	}

	res, err := s.cs.GetUserList(id)
	response := &pb.GetUserListResponse{
		Result: convertUser(res),
	}

	return response, err
}
