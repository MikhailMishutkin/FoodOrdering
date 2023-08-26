package stathandlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/statistics"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s StatisticService) TopProducts(
	ctx context.Context,
	in *statistics.TopProductsRequest,
) (
	*statistics.TopProductsResponse,
	error,
) {

	log.Println("TopProducts stathandlers was invoked")

	start := TimeAssert(in.StartDate)
	end := TimeAssert(in.EndDate)
	prType := int(in.GetProductType())

	res, err := s.SS.TopProducts(ctx, start, end, prType)
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "%v\n", err)
	}

	return &statistics.TopProductsResponse{
		Result: convertToRestProd(res),
	}, err
}
