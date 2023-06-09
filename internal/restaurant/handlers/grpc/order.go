package handlers

import (
	"context"
	"log"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
)

func (s *RestaurantService) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	log.Print("GetUpToDateOrderList was invoked")
	s.rSer.GetOrderList()
	return nil, nil
}
