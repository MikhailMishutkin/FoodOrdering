package natsstatistics

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/statistics"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
	"time"
)

func convertToProduct(pr []*restaurant.Product) (pt []*types.Product, err error) {
	for _, v := range pr {
		uuid, err := strconv.Atoi(v.Uuid)
		if err != nil {
			return nil, fmt.Errorf("can't convert product uuid in Profit stathandler: %v\n", err)
		}
		product := &types.Product{
			Uuid:      uuid,
			Name:      v.Name,
			Type:      int(v.Type.Number()),
			Price:     v.Price,
			CreatedAt: timeAssert(v.CreatedAt),
		}
		pt = append(pt, product)
	}
	return pt, err
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

func enumSelect(i int) statistics.StatisticsProductType {
	switch i {
	case 1:
		return statistics.StatisticsProductType_ST_PRODUCT_TYPE_SALAD
	case 2:
		return statistics.StatisticsProductType_ST_PRODUCT_TYPE_GARNISH
	case 3:
		return statistics.StatisticsProductType_ST_PRODUCT_TYPE_MEAT
	case 4:
		return statistics.StatisticsProductType_ST_PRODUCT_TYPE_SOUP
	case 5:
		return statistics.StatisticsProductType_ST_PRODUCT_TYPE_DRINK
	case 6:
		return statistics.StatisticsProductType_ST_PRODUCT_TYPE_DESSERT
	default:
		return statistics.StatisticsProductType_ST_PRODUCT_TYPE_UNSPECIFIED
	}

}
