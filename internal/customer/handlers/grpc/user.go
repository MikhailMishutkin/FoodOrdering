package handlers_customer

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func convertUser(res []*types.User) []*pb.User {
	var resPb []*pb.User

	for _, v := range res {
		id := strconv.Itoa(v.Uuid)
		ofId := strconv.Itoa(v.OfficeUuid)
		t := timestamppb.New(v.CreatedAt)
		pr := &pb.User{
			Uuid:       id,
			Name:       v.Name,
			OfficeUuid: ofId,
			OfficeName: v.OfficeName,
			CreatedAt:  t,
		}
		resPb = append(resPb, pr)
	}
	return resPb
}
