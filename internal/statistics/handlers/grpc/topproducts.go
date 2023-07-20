package stathandlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/statistics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
)

func (s StatisticService) TopProducts(
	ctx context.Context,
	in *statistics.TopProductsRequest,
) (
	*statistics.TopProductsResponse,
	error,
) {

	log.Println("TopProducts stathandlers was invoked")

	start := timeAssert(in.StartDate)
	end := timeAssert(in.EndDate)
	prType := int(in.GetProductType())

	res, err := s.SS.TopProducts(ctx, start, end, prType)
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "TopProducts went down witn error, cannot save products in db: %v\n", err)
	}

	return &statistics.TopProductsResponse{
		Result: convertToRestProd(res),
	}, err
}

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
