package handlers

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
	"time"
)

func timeAssert(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}

func enumSelect(i int) pb.ProductType {
	switch i {
	case 1:
		return pb.ProductType_PRODUCT_TYPE_SALAD
	case 2:
		return pb.ProductType_PRODUCT_TYPE_GARNISH
	case 3:
		return pb.ProductType_PRODUCT_TYPE_MEAT
	case 4:
		return pb.ProductType_PRODUCT_TYPE_SOUP
	case 5:
		return pb.ProductType_PRODUCT_TYPE_DRINK
	case 6:
		return pb.ProductType_PRODUCT_TYPE_DESSERT
	default:
		return pb.ProductType_PRODUCT_TYPE_UNSPECIFIED
	}

}

func convertProducts(res []*types.Product) []*pb.Product {
	var resPb []*pb.Product

	for _, v := range res {
		id := strconv.Itoa(v.Uuid)
		t := timestamppb.New(v.CreatedAt)
		pr := &pb.Product{
			Uuid:        id,
			Name:        v.Name,
			Description: v.Descript,
			Type:        enumSelect(v.Type),
			Weight:      int32(v.Weight),
			Price:       v.Price,
			CreatedAt:   t,
		}
		resPb = append(resPb, pr)
	}
	return resPb
}

func convertOffice(u *customer.Office) *types.Office {

	typesOffice := &types.Office{
		Uuid:      convStr(u.Uuid),
		Name:      u.Name,
		Address:   u.Address,
		CreatedAt: timeAssert(u.CreatedAt),
	}
	return typesOffice
}
func convertUser(u *customer.User) *types.User {

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
