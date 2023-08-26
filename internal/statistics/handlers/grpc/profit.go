package stathandlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/statistics"
)

func (s StatisticService) GetAmountOfProfit(
	ctx context.Context,
	in *statistics.GetAmountOfProfitRequest,
) (
	*statistics.GetAmountOfProfitResponse,
	error,
) {
	start := TimeAssert(in.StartDate)
	end := TimeAssert(in.EndDate)

	profit, err := s.SS.Profit(ctx, start, end)
	return &statistics.GetAmountOfProfitResponse{
		Profit: profit,
	}, err
}
