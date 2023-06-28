package serviceR

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/golang/protobuf/proto"
)

func (su *RestaurantUsecase) CreateProduct(b []byte) error {
	cpr := &restaurant.CreateProductRequest{}
	err := proto.Unmarshal(b, cpr)
	if err != nil {
		return fmt.Errorf("can't UNMARSHAL protodata to CreateProductRequest in service CreateProduct: %v\n", err)
	}
	p := extract(cpr)
	protoData, err := proto.Marshal(p)
	if err != nil {
		return fmt.Errorf("can't MARSHAL Product to protodata in service CreateProduct: %v\n", err)
	}
	err = su.repoR.CreateProduct(protoData)
	return err
}
func (su *RestaurantUsecase) GetProductList() ([]byte, error) {
	pl, err := su.repoR.GetProductList()
	return pl, err
}

func extract(cpr *restaurant.CreateProductRequest) *restaurant.Product {
	p := &restaurant.Product{
		Name:        cpr.Name,
		Description: cpr.Description,
		Type:        cpr.Type,
		Weight:      cpr.Weight,
		Price:       cpr.Price,
	}
	return p
}
