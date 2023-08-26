package natscustomer

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
)

type NatsPub struct {
	customer.UnimplementedOrderServiceServer
	ns NatsServicer
}

func NewNATS(ns NatsServicer) *NatsPub {
	return &NatsPub{
		ns: ns,
	}
}

type NatsServicer interface {
	CreateOrder(request *types.OrderRequest) error
}
