package repository

import (
	"sync"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/google/uuid"
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
