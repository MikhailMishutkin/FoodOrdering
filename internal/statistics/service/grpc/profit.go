package statservice

import (
	"context"
	"log"
	"time"
)

func (s *StatisticUsecase) Profit(
	ctx context.Context,
	start time.Time,
	end time.Time,
) (float64, error) {

	log.Println("Profit service was invoked")

	profit, err := s.sr.ProfitRepo(ctx, start, end)
	return profit, err
}
