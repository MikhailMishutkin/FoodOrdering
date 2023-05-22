package repository

import (
	"sync"

	"github.com/google/uuid"
	pb "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
)

var dataMap map[string]*pb.Product

func init() {
	dataMap = make(map[string]*pb.Product)
}
func RandomID() string {
	return uuid.New().String()
}

type RestaurantRepo struct {
	mutex   sync.RWMutex
	dataMap map[string]*pb.Product
}

func NewRestaurantRepo() *RestaurantRepo {
	return &RestaurantRepo{
		dataMap: dataMap,
	}
}

//TODO: перенести в отедльный файл Orders
func (r *RestaurantRepo) GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice) {

	return nil, nil
}
