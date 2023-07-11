package handlers_customer

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
)

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

func (s *CustomerService) GetActualMenu(ctx context.Context, in *pb.GetActualMenuRequest) (*pb.GetActualMenuResponse, error) {
	log.Println("GetActualMenu was invoked")

	t := time.Now()
	t1 := t.AddDate(0, 0, 1)
	ts := timestamppb.New(t1)

	fmt.Println(ts, s.client) //смотрим

	request, err := s.client.GetMenu(context.Background(), &restaurant.GetMenuRequest{
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

func convProductItem(slOI []*pb.OrderItem) (slTypOI []*types.OrderItem) {

	for _, v := range slOI {
		pId, err := strconv.Atoi(v.ProductUuid)
		if err != nil {
			fmt.Errorf("Can't conv user id in CreateOrder: %v\n", err)
		}
		typOI := &types.OrderItem{
			Count:       int(v.Count),
			ProductUuid: pId,
		}
		slTypOI = append(slTypOI, typOI)
	}
	return slTypOI
}

func convertProducts(res []*restaurant.Product) []*pb.Product {
	var resPb []*pb.Product

	for _, v := range res {
		pr := &pb.Product{
			Uuid:        v.Uuid,
			Name:        v.Name,
			Description: v.Description,
			Type:        enumSelect(int(v.Type.Number())),
			Weight:      v.Weight,
			Price:       v.Price,
			CreatedAt:   v.CreatedAt,
		}
		resPb = append(resPb, pr)
	}
	return resPb
}

func enumSelect(i int) pb.CustomerProductType {
	switch i {
	case 1:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_SALAD
	case 2:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_GARNISH
	case 3:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_MEAT
	case 4:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_SOUP
	case 5:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_DRINK
	case 6:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_DESSERT
	default:
		return pb.CustomerProductType_CUSTOMER_PRODUCT_TYPE_UNSPECIFIED
	}

}

func timeAssert(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}
