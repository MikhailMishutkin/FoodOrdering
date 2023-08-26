package stathandlers

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/statistics"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

func convertToRestProd(result []*types.StatProduct) (stResult []*statistics.Product) {
	for _, v := range result {
		p := &statistics.Product{
			Uuid:        strconv.Itoa(v.Uuid),
			Name:        v.Name,
			Count:       int64(v.Count),
			ProductType: enumSelect(v.Type),
		}
		stResult = append(stResult, p)
	}
	return stResult
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

func TimeAssert(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}
