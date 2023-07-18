package statservice

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"time"
)

type StatisticUsecase struct {
	sr StatisticRepositorier
}

func NewStatUsecase(sr StatisticRepositorier) *StatisticUsecase {
	return &StatisticUsecase{
		sr: sr,
	}
}

type StatisticRepositorier interface {
	ProfitRepo(time.Time, time.Time, []*types.Product) (float64, error)
	TopProductsRepo(context.Context, time.Time, time.Time, int) ([]*types.StatProduct, error)
}
