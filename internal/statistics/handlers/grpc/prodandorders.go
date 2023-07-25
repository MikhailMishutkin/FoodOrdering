package stathandlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *StatisticService) GetProduct() error {
	resProducts, err := s.ClientProduct.GetProductList(context.Background(), &restaurant.GetProductListRequest{})
	if err != nil {
		code := codes.Internal
		return status.Errorf(code, "GetProductList calling by Stat.Profit went down witn error, cannot save products in db: %v\n", err)
	}

	products, err := convertToTProduct(resProducts.Result)
	if err != nil {
		code := codes.Internal
		return status.Errorf(code, "convertToProduct went down witn error, cannot save products in db: %v\n", err)
	}

	err = s.SS.GetProducts(products)
	return err
}

func (s *StatisticService) GetOrders(order *types.OrderRequest) error {
	err := s.SS.GetOrders(order)
	return err
}
