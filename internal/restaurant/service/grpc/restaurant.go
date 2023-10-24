package serviceR

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"time"
)

type RestaurantUsecase struct {
	repoR RestaurantRepositorier
}

func NewRestaurantUsecase(rr RestaurantRepositorier) *RestaurantUsecase {
	return &RestaurantUsecase{
		repoR: rr,
	}
}

type RestaurantRepositorier interface {
	CreateProduct(product *types.Product) error
	GetProductList() ([]*types.Product, error)
	SelectProductByName(string, time.Time) (int, int, error)
	CreateDate(time.Time) (int, error)
	CreateMenu(int, int, int) error
	GetMenu(time.Time) (*types.Menu, error)
	GetTotalOrders(time.Time) ([]*types.OrderItem, error)
	GetOfficesList() ([]*types.OrderByOffice, error)
	GetOrdersByOffice(time.Time, int) ([]*types.OrderItem, error)
	ReceiveOrder(slOI []*types.OrderItem, userUuid int) error
}
