package gen

import (
	"fmt"

	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/contracts-v0.3.0/pkg/contracts/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

// generate products
func TypeSelector() {
	var m restaurant.ProductType
	fmt.Println(m)
	var p *restaurant.Product
	for {
		m++
		if m < 7 {
			p = NewProduct(m)
			(&repository.RestaurantRepo{}).CreateProduct(p)
		} else {
			break
		}
	}
}

func NewProduct(p restaurant.ProductType) *restaurant.Product {

	return &restaurant.Product{
		Uuid:        RandomID(),
		Name:        randomProductName(p),
		Description: randomDescription(p),
		Type:        p,
		Weight:      randomWeight(),
		Price:       randomPrice(),
		CreatedAt:   timestamppb.Now(),
	}
}
