package serviceR

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
)

func (su *RestaurantUsecase) CreateProduct(product *types.Product) error {
	err := su.repoR.CreateProduct(product)
	return err
}
func (su *RestaurantUsecase) GetProductList() ([]*types.Product, error) {
	pl, err := su.repoR.GetProductList()
	return pl, err
}

//func extract(cpr *restaurant.CreateProductRequest) *restaurant.Product {
//	p := &restaurant.Product{
//		Name:        cpr.Name,
//		Description: cpr.Description,
//		Type:        cpr.Type,
//		Weight:      cpr.Weight,
//		Price:       cpr.Price,
//	}
//	return p
//}
