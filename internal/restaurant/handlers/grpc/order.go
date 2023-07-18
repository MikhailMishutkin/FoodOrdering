package handlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
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
		&customer.GetOfficeListRequest{},
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
			&customer.GetUserListRequest{
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

func convertOrders(sl []*types.OrderItem) (slo []*pb.Order) {
	for _, v := range sl {
		o := &pb.Order{
			ProductId:   strconv.Itoa(v.ProductUuid),
			ProductName: v.ProductName,
			Count:       int64(v.Count),
		}
		slo = append(slo, o)
	}
	return slo
}

func convertOrdersByOffice(sl []*types.OrderByOffice) (tbo []*pb.OrdersByOffice) {

	for _, v := range sl {
		officeUuid := strconv.Itoa(v.OfficeUuid)
		obo := &pb.OrdersByOffice{
			OfficeUuid:    officeUuid,
			OfficeName:    v.OfficeName,
			OfficeAddress: v.OfficeAddress,
			Result:        convertOrders(v.Result),
		}
		tbo = append(tbo, obo)
	}
	return tbo
}
