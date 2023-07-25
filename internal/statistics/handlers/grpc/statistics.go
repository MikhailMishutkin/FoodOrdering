package stathandlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/statistics"
	"time"
)

type StatisticService struct {
	statistics.UnimplementedStatisticsServiceServer

	ClientProduct restaurant.ProductServiceClient
	SS            StatisticServicer
}

func NewStatService(clientProduct restaurant.ProductServiceClient, ss StatisticServicer) *StatisticService {
	return &StatisticService{
		ClientProduct: clientProduct,
		SS:            ss,
	}
}

type StatisticServicer interface {
	Profit(context.Context, time.Time, time.Time) (float64, error)
	TopProducts(context.Context, time.Time, time.Time, int) ([]*types.StatProduct, error)
	GetProducts([]*types.Product) error
	GetOrders(*types.OrderRequest) error
}
