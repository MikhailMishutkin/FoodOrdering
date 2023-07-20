package statservice

import "github.com/MikhailMishutkin/FoodOrdering/internal/types"

func (s *StatisticUsecase) GetProducts(products []*types.Product) error {
	err := s.sr.GetProductsRepo(products)
	return err
}

func (s *StatisticUsecase) GetOrders(orders *types.OrderRequest) error {
	err := s.sr.GetOrdersRepo(orders)
	return err
}
