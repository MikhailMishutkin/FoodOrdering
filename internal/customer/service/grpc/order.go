package service

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
)

func (cu *CustomerUsecase) CreateOrder(request *types.OrderRequest) error {
	log.Println("CreateOrder service was invoked")

	err := cu.nr.NatsPublisher(request)

	return err
}
