package handlers_customer

import (
	"context"
	"fmt"
	natscustomer "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/nats"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
)

func (s *CustomerService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	//res, err := s.cs.CreateOrder(in)
	a := &natscustomer.NatsPub{}
	err := a.NatsPublisher(in)
	return &pb.CreateOrderResponse{}, err
}

func (s *CustomerService) GetActualMenu(ctx context.Context, in *pb.GetActualMenuRequest) (*pb.GetActualMenuResponse, error) {
	log.Println("GetActualMenu was invoked")

	t := time.Now()
	t1 := t.AddDate(0, 0, 1)
	ts := timestamppb.New(t1)

	fmt.Println(ts, s.client) //смотрим

	result, err := s.client.GetMenu(context.Background(), &restaurant.GetMenuRequest{
		OnDate: ts,
	})
	if err != nil {
		log.Println("Can't get the menu from restaurant", err)
		return nil, err
	}

	amr, err := s.cs.GetActualMenu(result)
	if err != nil {
		return nil, err
	}
	return amr, nil
}
