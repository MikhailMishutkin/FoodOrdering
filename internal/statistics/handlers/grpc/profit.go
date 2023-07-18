package stathandlers

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/statistics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

func (s StatisticService) GetAmountOfProfit(
	ctx context.Context,
	in *statistics.GetAmountOfProfitRequest,
) (
	*statistics.GetAmountOfProfitResponse,
	error,
) {
	start := timeAssert(in.StartDate)
	end := timeAssert(in.EndDate)

	//call restaurant ProductList
	resProducts, err := s.client.GetProductList(context.Background(), &restaurant.GetProductListRequest{})
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "GetProductList calling by Stat.Profit went down witn error, cannot save products in db: %v\n", err)
	}

	products, err := convertToTProduct(resProducts.Result)
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "convertToProduct went down witn error, cannot save products in db: %v\n", err)
	}

	profit, err := s.ss.Profit(start, end, products)
	return &statistics.GetAmountOfProfitResponse{
		Profit: profit,
	}, err
}

func timeAssert(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}

func convertToTProduct(pr []*restaurant.Product) (pt []*types.Product, err error) {
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
