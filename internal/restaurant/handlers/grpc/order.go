package handlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	customer2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
)

func (s *RestaurantService) GetUpToDateOrderList(
	ctx context.Context,
	in *pb.GetUpToDateOrderListRequest,
) (*pb.GetUpToDateOrderListResponse, error) {
	log.Print("GetUpToDateOrderList was invoked")

	//get OfficeList
	requestOfficeList, err := s.OffClient.GetOfficeList(
		context.Background(),
		&customer2.GetOfficeListRequest{},
	)
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(
			code,
			"GetOfficeList went down witn error then calling by GetUpToDateOrderList: %v\n",
			err,
		)
	}

	sliceOfOffices := []*types.Office{}
	for _, v := range requestOfficeList.Result {
		office := convertOffice(v)
		sliceOfOffices = append(sliceOfOffices, office)
	}

	//get userlist
	sliceOfUsers := []*types.User{}
	for _, v := range sliceOfOffices {
		officeUuid := strconv.Itoa(v.Uuid)
		requestUserList, err := s.UsClient.GetUserList(
			context.Background(),
			&customer2.GetUserListRequest{
				OfficeUuid: officeUuid,
			},
		)
		if err != nil {
			code := codes.Internal
			return nil, status.Errorf(
				code,
				"GetUserList went down witn error then calling by GetUpToDateOrderList: %v\n",
				err,
			)
		}

		for _, v := range requestUserList.Result {
			user := convertUser(v)
			sliceOfUsers = append(sliceOfUsers, user)
		}

	}

	t, tbo, err := s.rSer.GetOrderList(sliceOfOffices, sliceOfUsers)

	return &pb.GetUpToDateOrderListResponse{
		TotalOrders:          convertOrders(t),
		TotalOrdersByCompany: convertOrdersByOffice(tbo),
	}, err
}
