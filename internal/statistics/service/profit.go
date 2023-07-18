package statservice

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

func (s *StatisticUsecase) Profit(
	start time.Time,
	end time.Time,
	products []*types.Product) (float64, error) {

	log.Println("Profit service was invoked")

	profit, err := s.sr.ProfitRepo(start, end, products)
	return profit, err
}
