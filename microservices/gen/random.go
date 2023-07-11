package gen

import (
	"math/rand"
	"time"

	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomID() string {
	return uuid.New().String()
}

func randomOffice() (string, string) {
	ofname := randomStringFromSet("Ателье", "Магазин Бум", "Офис Гранд")
	ofAdr := randomStringFromSet("Гончарова 22, оф. 3", "К. Маркса 5а, оф. 17", "ул. Мира 7, оф. 36")
	return ofname, ofAdr
}

func randomProductName(t restaurant.ProductType) string {

	switch t {
	case restaurant.ProductType_PRODUCT_TYPE_SALAD:
		return randomStringFromSet("Оливье", "Цезарь", "Винигрет")
	case restaurant.ProductType_PRODUCT_TYPE_GARNISH:
		return randomStringFromSet("Гречка", "Рис", "Пюре")
	case restaurant.ProductType_PRODUCT_TYPE_MEAT:
		return randomStringFromSet("Котлета", "Сарделька", "Сосиска")
	case restaurant.ProductType_PRODUCT_TYPE_SOUP:
		return randomStringFromSet("Борщ", "Суп с фрикадельками", "Харчо")
	case restaurant.ProductType_PRODUCT_TYPE_DRINK:
		return randomStringFromSet("Компот", "Морс", "Квас")
	case restaurant.ProductType_PRODUCT_TYPE_DESSERT:
		return randomStringFromSet("Картошка", "Медовик", "Эклер")
	default:
		return "Тип продукта не определён"
	}
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomDescription(t restaurant.ProductType) string {
	switch t {
	case restaurant.ProductType_PRODUCT_TYPE_SALAD:
		return randomStringFromSet("Хорошо сочетается с мясом", "Витаминный заряд", "Легкий и вкусный")
	case restaurant.ProductType_PRODUCT_TYPE_GARNISH:
		return randomStringFromSet("Отварное", "На молоке")
	case restaurant.ProductType_PRODUCT_TYPE_MEAT:
		return randomStringFromSet("Куриная", "Говяжья", "Свинная")
	case restaurant.ProductType_PRODUCT_TYPE_SOUP:
		return randomStringFromSet("Со сметаной", "Острый", "Диетический")
	case restaurant.ProductType_PRODUCT_TYPE_DRINK:
		return randomStringFromSet("Из черешни", "Ягодный", "Домашний")
	case restaurant.ProductType_PRODUCT_TYPE_DESSERT:
		return randomStringFromSet("Белковый крем", "Масляный крем", "Диетический без сахара")
	default:
		return "Тип продукта не определён"
	}
}

func randomWeight() int32 {
	var w, min, max int32
	min = 100
	max = 250
	w = int32(randomInt(int(min), int(max)))
	return w
}

func randomPrice() float64 {
	var p, min, max float64
	min = 100
	max = 151
	p = randomFloat64(min, max)
	return p

}

func randomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
