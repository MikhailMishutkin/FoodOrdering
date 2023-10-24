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

	//order := &types.OrderRequest{}

	//sub, err := n.Conn.SubscribeSync("order")
	//if err != nil {
	//	log.Printf("subscribeSync error: %v\n: ", err)
	//}

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

	//for {
	//	t := repository.DateConv(time.Now())
	//	t1 := t.AddDate(0, 0, 1)
	//	t2 := t1.Add(11 * time.Hour)
	//	t3 := t1.Add(21 * time.Hour)
	//
	//	msg, err := sub.NextMsgWithContext(context.Background())
	//	if err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//	if msg.Subject == "order" {
	//		err = json.Unmarshal(msg.Data, order)
	//		err = n.Jm.DataSaveService(order, sliceOfOffices, sliceOfUsers)
	//		if err != nil {
	//			log.Println(err)
	//		}
	//	}
	//	if time.Now().UnixNano() >= t2.UnixNano() && time.Now().UnixNano() < t3.UnixNano() {
	//
	//	} else {
	//		continue
	//	}
	//
	//}
	return err
}

//menu, err := n.repo.GetMenu(t1)
//if err != nil {
//	fmt.Errorf("restaurant GetMenu error: %v\n", err)
//}
