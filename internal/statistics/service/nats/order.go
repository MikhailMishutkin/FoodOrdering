package natsstatservice

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
)

func (r *StatNATSService) DataSaveService(order *types.OrderRequest, products []*types.Product) error {

	for _, v := range products {
		err := r.nsr.SaveProduct(v)
		if err != nil {
			return err
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
		//TODO
		res2B, _ := json.Marshal(v)
		fmt.Println(string(res2B))

		if v.Count != 0 && v.ProductUuid != 0 {
			err := r.nsr.SaveOrderStat(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
