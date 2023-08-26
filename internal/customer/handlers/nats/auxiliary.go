package natscustomer

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"strconv"
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
