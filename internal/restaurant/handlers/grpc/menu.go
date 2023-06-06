package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *RestaurantService) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	log.Print("CreateMenu was invoked")

	res, err := s.repoR.CreateMenu()
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "repo.CreateMenu went down witn error, cannot create menu: %v/n ", err)
	}

	file, err := os.OpenFile("menu.json", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	j, _ := json.Marshal(res)
	file.Write(j)

	resp := &pb.CreateMenuResponse{}

	return resp, nil
}

func (s *RestaurantService) GetMenu(ctx context.Context, in *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	log.Print("GetMenu was invoked")
	ts := in.OnDate
	t := ts.AsTime()
	m := &pb.Menu{}
	m, err := s.repoR.GetMenu(t)
	if err != nil {
		return nil, err
	}
	//go repository.NatsPublisher()
	resp := &pb.GetMenuResponse{}
	resp.Menu = m

	return resp, nil
}
