package serviceR

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"time"
)

type RestaurantUsecase struct {
	repoR RestaurantRepository
}

func NewRestaurantUsecace(rr RestaurantRepository) *RestaurantUsecase {
	return &RestaurantUsecase{
		repoR: rr,
	}
}

type RestaurantRepository interface {
	CreateProduct(product *types.Product) error
	GetProductList() ([]*types.Product, error)
	CreateMenu(mc *types.MenuCreate) error
	GetMenu(time.Time) (*types.Menu, error)
	GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice)
}
