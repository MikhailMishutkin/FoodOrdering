package serviceR

import (
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

func (su *RestaurantUsecase) CreateProduct(p *pb.Product) error {
	err := su.repoR.CreateProduct(p)
	return err
}
func (su *RestaurantUsecase) GetProductList() (*pb.GetProductListResponse, error) {
	return nil, nil
}
