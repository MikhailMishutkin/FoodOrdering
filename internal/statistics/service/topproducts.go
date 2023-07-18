package statservice

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

func (s StatisticUsecase) TopProducts(
	ctx context.Context,
	start time.Time,
	end time.Time,
	prType int,
) (
	products []*types.StatProduct,
	err error,
) {
	log.Println("TopProducts statservice was invoked")

	products, err = s.sr.TopProductsRepo(ctx, start, end, prType)

	return products, err
}
