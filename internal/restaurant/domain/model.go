package domain

import (
	"reflect"

	"github.com/setrofim/enum"
)

type Menu struct {
	uuid              string
	on_date           string
	opening_record_at string
	closing_record_at string
	salads            restaurantProduct
	garnishes         restaurantProduct
	meats             restaurantProduct
	soups             restaurantProduct
	drinks            restaurantProduct
	desserts          restaurantProduct
	created_at        string
}

type restaurantProduct struct {
	uuid        string
	name        string
	description string
	type_       restaurantProductType
	weight      int
	price       float64
	created_at  string
}

type restaurantProductType string

func (restaurantProductType) PRODUCT_TYPE_UNSPECIFIED() restaurantProductType {
	return restaurantProductType(0)
}
func (restaurantProductType) PRODUCT_TYPE_SALAD() restaurantProductType {
	return restaurantProductType(1)
}
func (restaurantProductType) PRODUCT_TYPE_GARNISH() restaurantProductType {
	return restaurantProductType(2)
}
func (restaurantProductType) PRODUCT_TYPE_MEAT() restaurantProductType {
	return restaurantProductType(3)
}
func (restaurantProductType) PRODUCT_TYPE_SOUP() restaurantProductType {
	return restaurantProductType(4)
}
func (restaurantProductType) PRODUCT_TYPE_DRINK() restaurantProductType {
	return restaurantProductType(5)
}
func (restaurantProductType) PRODUCT_TYPE_DESSERT() restaurantProductType {
	return restaurantProductType(6)
}

func (c restaurantProductType) String() string {
	return enum.StringInt(c, reflect.TypeOf(c))
}

func (c *restaurantProductType) Parse(s string) error {
	enumVal, err := enum.ParseInt(reflect.TypeOf(c), s, true, false)
	if enumVal != nil {
		*c = enumVal.(restaurantProductType) // If no error, type assert to Color and set c
	}
	return err
}
