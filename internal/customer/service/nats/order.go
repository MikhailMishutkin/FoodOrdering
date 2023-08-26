package natscustomerservice

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
)

func (nu *CustomerNatsUsecase) CreateOrder(request *types.OrderRequest) error {
	log.Println("CreateOrder service was invoked")

	err := nu.nr.NatsPublisher(request)

	return err
}
