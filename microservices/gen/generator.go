package gen

import (
	"fmt"

	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	res "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

// generate products
func TypeSelector() {
	var m res.ProductType
	fmt.Println(m)
	var p *res.Product
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

func NewProduct(p res.ProductType) *res.Product {

	return &res.Product{
		Uuid:        RandomID(),
		Name:        randomProductName(p),
		Description: randomDescription(p),
		Type:        p,
		Weight:      randomWeight(),
		Price:       randomPrice(),
		CreatedAt:   timestamppb.Now(),
	}
}
