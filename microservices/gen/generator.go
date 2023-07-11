package gen

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	res "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"log"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

// generate products
func TypeSelector() {
	var m res.ProductType
	p := &types.Product{}
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	for {
		m++
		if m < 7 {
			p = NewProduct(m)
			repo := &repository.RestaurantRepo{
				DB: db,
			}
			repo.CreateProduct(p)

		} else {
			break
		}
	}

}

func NewProduct(t res.ProductType) *types.Product {

	p := &types.Product{
		//Uuid:        RandomID(),
		Name:     randomProductName(t),
		Descript: randomDescription(t),
		Type:     enumSelect(t),
		Weight:   int(randomWeight()),
		Price:    randomPrice(),
	}

	return p
}

func NewOffice() *types.Office {
	n, a := randomOffice()
	return &types.Office{
		Name:    n,
		Address: a,
	}
}

func enumSelect(t res.ProductType) int {
	switch t {
	case res.ProductType_PRODUCT_TYPE_SALAD:
		return 1
	case res.ProductType_PRODUCT_TYPE_GARNISH:
		return 2
	case res.ProductType_PRODUCT_TYPE_MEAT:
		return 3
	case res.ProductType_PRODUCT_TYPE_SOUP:
		return 4
	case res.ProductType_PRODUCT_TYPE_DRINK:
		return 5
	case res.ProductType_PRODUCT_TYPE_DESSERT:
		return 6
	default:
		return 0
	}

}

//func OfficeGen() {
//	var o *cus.Office
//	o = NewOffice()
//	(&handlers_customer.CustomerService{}).CreateOffice(o)
//}
