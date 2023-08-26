package statrepository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *StatRepo) ReceiveOrdersRepo(order *types.OrderRequest) error {
	log.Println("GetOrder (stat_repo statistics) was invoked")

	resProducts, err := s.ClientProduct.GetProductList(context.Background(), &restaurant.GetProductListRequest{})
	if err != nil {
		code := codes.Internal
		return status.Errorf(code, "GetProductList calling by Stat.Profit went down witn error, cannot save products in db: %v\n", err)
	}

	products, err := convertToTProduct(resProducts.Result)
	if err != nil {
		code := codes.Internal
		return status.Errorf(code, "convertToProduct went down witn error, cannot save products in db: %v\n", err)
	}

	const q = `
		INSERT INTO product (uuid, name, description, type_id, weight, price, created_at) 
		VALUES(:uuid, :name, :description, :type_id, :weight, :price, :created_at) ON CONFLICT DO NOTHING`
	for _, v := range products {
		_, err := s.DB.NamedExec(q, v)
		if err != nil {

		}
	}
	var slOI []*types.OrderItem

	slOI = append(slOI, order.Salads...)
	slOI = append(slOI, order.Garnishes...)
	slOI = append(slOI, order.Meats...)
	slOI = append(slOI, order.Soups...)
	slOI = append(slOI, order.Drinks...)
	slOI = append(slOI, order.Desserts...)

	for _, v := range slOI {
		res2B, _ := json.Marshal(v)
		fmt.Println(string(res2B))

		if v.Count != 0 && v.ProductUuid != 0 {
			_, err := s.DB.NamedExec(
				"INSERT INTO orders (product_id, count) VALUES (:product_id, :count)",
				v,
			)
			if err != nil {
				return fmt.Errorf("Can't INSERT order into db stat: %v\n", err)
			}
		}
	}

	return nil

}
