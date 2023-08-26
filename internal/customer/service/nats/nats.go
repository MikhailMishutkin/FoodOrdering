package natscustomerservice

import "github.com/MikhailMishutkin/FoodOrdering/internal/types"

type CustomerNatsUsecase struct {
	nr NatsCustomerRepositorier
}

func NewCNU(nr NatsCustomerRepositorier) *CustomerNatsUsecase {
	return &CustomerNatsUsecase{
		nr: nr,
	}
}

type NatsCustomerRepositorier interface {
	NatsPublisher(order *types.OrderRequest) error
}
