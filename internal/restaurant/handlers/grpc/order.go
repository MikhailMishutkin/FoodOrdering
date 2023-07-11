package handlers

import (
	"context"
	"log"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

func (s *RestaurantService) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	log.Print("GetUpToDateOrderList was invoked")
	a, _ := s.rSer.GetOrderList()
	return &pb.GetUpToDateOrderListResponse{
		TotalOrders: a,
	}, nil
}
