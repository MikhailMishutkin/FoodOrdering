package natsrestaurant

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	customer2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
	"time"
)

func convertOffice(u *customer2.Office) *types.Office {

	typesOffice := &types.Office{
		Uuid:      convStr(u.Uuid),
		Name:      u.Name,
		Address:   u.Address,
		CreatedAt: timeAssert(u.CreatedAt),
	}
	return typesOffice
}
func convertUser(u *customer2.User) *types.User {

	typesUser := &types.User{
		Uuid:       convStr(u.Uuid),
		Name:       u.Name,
		OfficeUuid: convStr(u.OfficeUuid),
		OfficeName: u.OfficeName,
		CreatedAt:  timeAssert(u.CreatedAt),
	}
	return typesUser
}

func convStr(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("can't convert string to int in GetUpToDateOrderList handler: %v\n", err)
	}
	return i
}

func timeAssert(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}
