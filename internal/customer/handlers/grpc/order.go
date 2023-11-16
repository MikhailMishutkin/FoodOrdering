package handlerscustomer

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	restaurant2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
	"time"
)

func (s *CustomerService) GetActualMenu(ctx context.Context, in *pb.GetActualMenuRequest) (*pb.GetActualMenuResponse, error) {
	log.Println("GetActualMenu was invoked")

	t := time.Now()
	t1 := t.AddDate(0, 0, 1)
	ts := timestamppb.New(t1)

	fmt.Println(ts, s.client) //смотрим

	request, err := s.client.GetMenu(context.Background(), &restaurant2.GetMenuRequest{
		OnDate: ts,
	})
	if err != nil {
		log.Println("Can't get the menu from restaurant", err)
		return nil, err
	}

	result := &pb.GetActualMenuResponse{
		Salads:    convertProducts(request.Menu.Salads),
		Garnishes: convertProducts(request.Menu.Garnishes),
		Meats:     convertProducts(request.Menu.Meats),
		Soups:     convertProducts(request.Menu.Soups),
		Drinks:    convertProducts(request.Menu.Drinks),
		Desserts:  convertProducts(request.Menu.Desserts),
	}

	return result, err
}

func (s *CustomerService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
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

	err = s.cs.CreateOrder(req)

	return &pb.CreateOrderResponse{}, err

}
