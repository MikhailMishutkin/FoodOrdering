package handlerscustomer

import (
	"context"
	"fmt"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	restaurant2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
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
