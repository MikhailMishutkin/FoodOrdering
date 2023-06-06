package gen

import (
	"log"

	handlers_customer "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/grpc"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	cus "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	res "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

// generate products
func TypeSelector() {
	var m res.ProductType
	log.Println(m)
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

func OfficeGen() {
	var o *cus.Office
	o = NewOffice()
	(&handlers_customer.CustomerService{}).CreateOffice(o)
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

func NewOffice() *cus.Office {
	n, a := randomOffice()
	return &cus.Office{
		Uuid:      RandomID(),
		Name:      n,
		Address:   a,
		CreatedAt: timestamppb.Now(),
	}
}
