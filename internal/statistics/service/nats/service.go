package natsstatservice

import "github.com/MikhailMishutkin/FoodOrdering/internal/types"

type StatNATSService struct {
	nsr NATSStatRepositorier
}

func NewSNService(nsr NATSStatRepositorier) *StatNATSService {
	return &StatNATSService{
		nsr: nsr,
	}
}

type NATSStatRepositorier interface {
	SaveOrderStat(*types.OrderItem) error
	SaveProduct(*types.Product) error
}
