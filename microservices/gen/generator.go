package gen

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	cus "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	res "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	//"google.golang.org/protobuf/types/known/timestamppb"
)

// generate products
func TypeSelector() {
	var m res.ProductType
	//log.Println(m)
	var p []byte
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

//func OfficeGen() {
//	var o *cus.Office
//	o = NewOffice()
//	(&handlers_customer.CustomerService{}).CreateOffice(o)
//}

func NewProduct(t res.ProductType) []byte {

	p := &res.Product{
		//Uuid:        RandomID(),
		Name:        randomProductName(t),
		Description: randomDescription(t),
		Type:        t,
		Weight:      randomWeight(),
		Price:       randomPrice(),
		CreatedAt:   timestamppb.Now(),
	}
	data, err := proto.Marshal(p)
	if err != nil {
		fmt.Errorf("can't MARSHAL data in generator NewProduct: %v\n", err)

	}
	return data
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
