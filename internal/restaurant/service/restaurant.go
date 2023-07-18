package serviceR

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"time"
)

type RestaurantUsecase struct {
	repoR RestaurantRepositorier
}

func NewRestaurantUsecace(rr RestaurantRepositorier) *RestaurantUsecase {
	return &RestaurantUsecase{
		repoR: rr,
	}
}

type RestaurantRepositorier interface {
	CreateProduct(product *types.Product) error
	GetProductList() ([]*types.Product, error)
	CreateMenu(mc *types.MenuCreate) error
	GetMenu(time.Time) (*types.Menu, error)
	GetOrderList(time.Time, []*types.Office, []*types.User) ([]*types.OrderItem, []*types.OrderByOffice, error)
	GetOrder(order *types.OrderRequest) error
}
