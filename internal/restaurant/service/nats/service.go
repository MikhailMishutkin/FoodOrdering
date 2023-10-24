package natsrestservice

import "github.com/MikhailMishutkin/FoodOrdering/internal/types"

type RestNATSService struct {
	nrr NATSRestRepositorier
}

func NewRNService(nrr NATSRestRepositorier) *RestNATSService {
	return &RestNATSService{
		nrr: nrr,
	}
}

type NATSRestRepositorier interface {
	ReceiveOrder([]*types.OrderItem, int) error
	SaveOfficeList([]*types.Office) error
	SaveUserList([]*types.User) error
}
