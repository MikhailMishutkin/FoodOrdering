package statrepository

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

func (s StatRepo) ProfitRepo(
	time time.Time,
	time2 time.Time,
	products []*types.Product,
) (profit float64, err error) {

	log.Println("ProfitRepo was invoked")

	return profit, err
}
