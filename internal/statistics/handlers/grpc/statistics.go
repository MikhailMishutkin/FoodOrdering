package stathandlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/statistics"
	"time"
)

type StatisticService struct {
	statistics.UnimplementedStatisticsServiceServer
	SS StatisticServicer
}

func NewStatService(ss StatisticServicer) *StatisticService {
	return &StatisticService{
		SS: ss,
	}
}

type StatisticServicer interface {
	Profit(context.Context, time.Time, time.Time) (float64, error)
	TopProducts(context.Context, time.Time, time.Time, int) ([]*types.StatProduct, error)
}
