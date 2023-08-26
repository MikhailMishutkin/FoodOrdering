package natscustomer

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"log"
	"strconv"
)

func (n *NatsPub) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Println("CreateOrder was invoked")
	usId, err := strconv.Atoi(in.UserUuid)
	if err != nil {
		return nil, fmt.Errorf("Can't conv user id in CreateOrder: %v\n", err)
	}

	req := &types.OrderRequest{
		UserUuid:  usId,
		Salads:    convProductItem(in.Salads),
		Garnishes: convProductItem(in.Garnishes),
		Meats:     convProductItem(in.Meats),
		Soups:     convProductItem(in.Soups),
		Drinks:    convProductItem(in.Drinks),
		Desserts:  convProductItem(in.Desserts),
	}

	err = n.ns.CreateOrder(req)

	return &pb.CreateOrderResponse{}, err

}
