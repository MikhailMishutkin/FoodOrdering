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

	client restaurant.ProductServiceClient
	ss     StatisticServicer
}

func NewStatService(client restaurant.ProductServiceClient, ss StatisticServicer) *StatisticService {
	return &StatisticService{
		client: client,
		ss:     ss,
	}
}

type StatisticServicer interface {
	Profit(time.Time, time.Time, []*types.Product) (float64, error)
	TopProducts(context.Context, time.Time, time.Time, int) ([]*types.StatProduct, error)
}
