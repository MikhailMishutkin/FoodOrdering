package statrepository

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
	"time"
)

type TopProduct struct {
	ID int

	Count  int
	TypeID int
}

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

	const q = `SELECT orders.product_id, product.name, SUM(orders.count), product.type_id
						FROM orders, product
							WHERE  orders.on_date >=$1::date
								AND    orders.on_date <=$2::date
								AND orders.product_id = product.uuid
								AND product.type_id = $3
							GROUP BY orders.product_id, product.name, orders.count, product.type_id
							ORDER BY count(*) DESC
							LIMIT 2`

	err = s.DB.SelectContext(ctx, &products, q, start, end, prType)
	if err != nil {
		return nil, fmt.Errorf("Something wrong with select product count and type from db in TopProduct stat: %v\n", err)
	}

	return products, err
}
