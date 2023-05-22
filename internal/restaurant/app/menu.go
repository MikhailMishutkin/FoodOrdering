package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	pb "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	log.Print("CreateMenu was invoked")

	res, err := s.repo.CreateMenu()
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "repo.CreateMenu went down witn error, cannot create menu: ", err)
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

func (s *Server) GetMenu(ctx context.Context, in *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	log.Print("GetMenu was invoked")

	ts := in.OnDate
	fmt.Println("время из запроса постман:", ts)
	t := ts.AsTime()
	fmt.Println("время преобразованное в time:", t)
	m := s.repo.GetMenu(t)
	go repository.NatsPublisher()
	resp := &pb.GetMenuResponse{}
	resp.Menu = m

	return resp, nil
}
