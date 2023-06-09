package serviceR

import (
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"time"
)

func (su *RestaurantUsecase) CreateMenu() (*pb.Menu, error) {
	m, err := su.repoR.CreateMenu()
	return m, err
}

func (su *RestaurantUsecase) GetMenu(t time.Time) (*pb.Menu, error) {
	m, err := su.repoR.GetMenu(t)
	return m, err
}

func (su *RestaurantUsecase) GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice) {
	return nil, nil
}
