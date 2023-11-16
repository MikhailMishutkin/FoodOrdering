package statrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
)

func (s *StatRepo) SaveOrderStat(order *types.OrderItem) error {
	log.Println("GetOrder (stat_repo statistics) was invoked")

	_, err := s.DB.NamedExec(
		"INSERT INTO orders (product_id, count) VALUES (:product_id, :count)",
		order,
	)
	if err != nil {
		return fmt.Errorf("Can't INSERT order into db stat: %v\n", err)
	}

	return err

}

func (s *StatRepo) SaveProduct(pr *types.Product) error {
	const q = `
		INSERT INTO product (uuid, name, description, type_id, weight, price, created_at) 
		VALUES(:uuid, :name, :description, :type_id, :weight, :price, :created_at) ON CONFLICT DO NOTHING`
	_, err := s.DB.NamedExec(q, pr)
	if err != nil {
		return fmt.Errorf("can't save product to db: %v\n", err)
	}
	return err
}
