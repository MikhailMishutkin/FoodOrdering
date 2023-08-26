package handlers

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	customer2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
	"time"
)

func timeAssert(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}

func enumSelect(i int) restaurant.ProductType {
	switch i {
	case 1:
		return restaurant.ProductType_PRODUCT_TYPE_SALAD
	case 2:
		return restaurant.ProductType_PRODUCT_TYPE_GARNISH
	case 3:
		return restaurant.ProductType_PRODUCT_TYPE_MEAT
	case 4:
		return restaurant.ProductType_PRODUCT_TYPE_SOUP
	case 5:
		return restaurant.ProductType_PRODUCT_TYPE_DRINK
	case 6:
		return restaurant.ProductType_PRODUCT_TYPE_DESSERT
	default:
		return restaurant.ProductType_PRODUCT_TYPE_UNSPECIFIED
	}

}

func convertProducts(res []*types.Product) []*restaurant.Product {
	var resPb []*restaurant.Product

	for _, v := range res {
		id := strconv.Itoa(v.Uuid)
		t := timestamppb.New(v.CreatedAt)
		pr := &restaurant.Product{
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

func convertOrders(sl []*types.OrderItem) (slo []*restaurant.Order) {
	for _, v := range sl {
		o := &restaurant.Order{
			ProductId:   strconv.Itoa(v.ProductUuid),
			ProductName: v.ProductName,
			Count:       int64(v.Count),
		}
		slo = append(slo, o)
	}
	return slo
}

func convertOrdersByOffice(sl []*types.OrderByOffice) (tbo []*restaurant.OrdersByOffice) {

	for _, v := range sl {
		officeUuid := strconv.Itoa(v.OfficeUuid)
		obo := &restaurant.OrdersByOffice{
			OfficeUuid:    officeUuid,
			OfficeName:    v.OfficeName,
			OfficeAddress: v.OfficeAddress,
			Result:        convertOrders(v.Result),
		}
		tbo = append(tbo, obo)
	}
	return tbo
}
