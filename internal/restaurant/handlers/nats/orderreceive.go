package natsrestaurant

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	customer2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
)

func (n *NatsSub) OrderReceive(order *types.OrderRequest) error {
	log.Println("OrderReceive was invoked")

	//get OfficeList
	requestOfficeList, err := n.OffClient.GetOfficeList(
		context.Background(),
		&customer2.GetOfficeListRequest{},
	)
	if err != nil {
		code := codes.Internal
		return status.Errorf(
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
		requestUserList, err := n.UsClient.GetUserList(
			context.Background(),
			&customer2.GetUserListRequest{
				OfficeUuid: officeUuid,
			},
		)
		if err != nil {
			code := codes.Internal
			return status.Errorf(
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

	err = n.Jm.DataSaveService(order, sliceOfOffices, sliceOfUsers)
	if err != nil {
		log.Println(err)
	}

	return err
}
