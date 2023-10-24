package handlers

import (
	"context"
	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func (s *RestaurantService) GetUpToDateOrderList(
	ctx context.Context,
	in *pb.GetUpToDateOrderListRequest,
) (*pb.GetUpToDateOrderListResponse, error) {
	log.Print("GetUpToDateOrderList was invoked")

	t, tbo, err := s.rSer.GetOrderList()

	return &pb.GetUpToDateOrderListResponse{
		TotalOrders:          convertOrders(t),
		TotalOrdersByCompany: convertOrdersByOffice(tbo),
	}, err
}
