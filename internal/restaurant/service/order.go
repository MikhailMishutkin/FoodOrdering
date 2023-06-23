package serviceR

import pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"

func (su *RestaurantUsecase) GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice) {
	a, _ := su.repoR.GetOrderList()
	return a, nil
}
