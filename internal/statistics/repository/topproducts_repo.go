package statrepository

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

func (s *StatRepo) TopProductsRepo(
	ctx context.Context,
	start time.Time,
	end time.Time,
	prType int,
) (
	products []*types.StatProduct,
	err error,
) {
	log.Println("TopProductsRepo was invoked")

	return products, err
}
