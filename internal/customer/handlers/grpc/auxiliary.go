package handlerscustomer

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	restaurant2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

func convProductItem(slOI []*pb.OrderItem) (slTypOI []*types.OrderItem) {

	for _, v := range slOI {
		pId, err := strconv.Atoi(v.ProductUuid)
		if err != nil {
			fmt.Errorf("Can't conv user id in CreateOrder: %v\n", err)
		}
		typOI := &types.OrderItem{
			Count:       int(v.Count),
			ProductUuid: pId,
		}
		slTypOI = append(slTypOI, typOI)
	}
	return slTypOI
}

func convertProducts(res []*restaurant2.Product) []*pb.Product {
	var resPb []*pb.Product

	for _, v := range res {
		pr := &pb.Product{
			Uuid:        v.Uuid,
			Name:        v.Name,
			Description: v.Description,
			Type:        enumSelect(int(v.Type.Number())),
			Weight:      v.Weight,
			Price:       v.Price,
			CreatedAt:   v.CreatedAt,
		}
		resPb = append(resPb, pr)
	}
	return resPb
}

func enumSelect(i int) pb.CustomerProductType {
	switch i {
	case 1:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_SALAD
	case 2:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_GARNISH
	case 3:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_MEAT
	case 4:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_SOUP
	case 5:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_DRINK
	case 6:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_DESSERT
	default:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_UNSPECIFIED
	}

}

func timeAssert(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}

func convertOffice(res []*types.Office) []*pb.Office {
	var resPb []*pb.Office

	for _, v := range res {
		id := strconv.Itoa(v.Uuid)
		t := timestamppb.New(v.CreatedAt)
		pr := &pb.Office{
			Uuid:      id,
			Name:      v.Name,
			Address:   v.Address,
			CreatedAt: t,
		}
		resPb = append(resPb, pr)
	}
	return resPb
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
