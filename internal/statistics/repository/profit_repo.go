package statrepository

import (
	"context"
	"log"
	"time"
)

func (s StatRepo) ProfitRepo(
	ctx context.Context,
	time time.Time,
	time2 time.Time,
) (profit float64, err error) {

	log.Println("ProfitRepo was invoked")

	return profit, err
}
