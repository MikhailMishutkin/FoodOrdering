package handlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
)

func (s *RestaurantService) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	log.Print("CreateMenu was invoked")

	mc := &types.MenuCreate{
		OnDate:    timeAssert(in.OnDate),
		OpenAt:    timeAssert(in.OpeningRecordAt),
		ClosedAt:  timeAssert(in.ClosingRecordAt),
		Salads:    in.Salads,
		Garnishes: in.Garnishes,
		Meats:     in.Meats,
		Soups:     in.Soups,
		Drinks:    in.Drinks,
		Desserts:  in.Desserts,
	}
	err := s.rSer.CreateMenu(mc)
	log.Println(err)
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "repo.CreateMenu went down witn error, cannot create menu: %v/n ", err)
	}

	resp := &pb.CreateMenuResponse{}

	return resp, nil
}

func (s *RestaurantService) GetMenu(ctx context.Context, in *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	log.Print("GetMenu was invoked")
	t := timeAssert(in.OnDate)
	//t := ts.AsTime()

	m, err := s.rSer.GetMenu(t)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetMenuResponse{}
	rm := &pb.Menu{
		Uuid:            strconv.Itoa(m.Uuid),
		OnDate:          in.OnDate,
		OpeningRecordAt: timestamppb.New(m.OpenAt),
		ClosingRecordAt: timestamppb.New(m.ClosedAt),
		Salads:          convertProducts(m.Salads),
		Garnishes:       convertProducts(m.Garnishes),
		Meats:           convertProducts(m.Meats),
		Soups:           convertProducts(m.Soups),
		Drinks:          convertProducts(m.Drinks),
		Desserts:        convertProducts(m.Desserts),
		CreatedAt:       timestamppb.New(m.CreatedAt),
	}

	resp.Menu = rm

	return resp, nil
}
