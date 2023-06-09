package serviceR

import (
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
	CreateProduct(*pb.Product) error
	GetProductList() (*pb.GetProductListResponse, error)
	CreateMenu() (*pb.Menu, error)
	GetMenu(time.Time) (*pb.Menu, error)
	GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice)
}
