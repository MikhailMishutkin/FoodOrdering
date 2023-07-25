package statrepository

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

func (s StatRepo) ProfitRepo(
	ctx context.Context,
	start time.Time,
	end time.Time,
) (profit float64, err error) {

	log.Println("ProfitRepo was invoked")

	const q = `SELECT orders.count, product.price
FROM orders, product
     WHERE  orders.on_date >=$1::date
       AND    orders.on_date   <=$2::date
        AND orders.product_id = product.uuid`
	var p []*types.Profit
	err = s.DB.SelectContext(ctx, &p, q, start, end)
	if err != nil {
		return 0, fmt.Errorf("Something wrong with select product count and price from db in profit stat: %v\n", err)
	}
	for _, v := range p {
		profit = profit + (float64(v.Count) * v.Price)
	}

	return profit, err
}
