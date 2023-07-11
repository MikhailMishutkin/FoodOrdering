package handlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// generated 12 products by package gen if name == ""
func (s *RestaurantService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	log.Print("CreateProduct was invoked")

	if in.Name == "" {
		gen.TypeSelector()
		gen.TypeSelector()
	} else {
		p := &types.Product{
			Name:     in.Name,
			Descript: in.Description,
			Type:     int(in.Type.Number()),
			Weight:   int(in.Weight),
			Price:    in.Price,
		}
		err := s.rSer.CreateProduct(p)
		if err != nil {
			code := codes.Internal
			return nil, status.Errorf(code, "repo.CreateProduct went down witn error, cannot save product in db: %v\n", err)
		}
	}

	resp := &pb.CreateProductResponse{}

	return resp, nil
}

func (s *RestaurantService) GetProductList(ctx context.Context, in *pb.GetProductListRequest) (*pb.GetProductListResponse, error) {
	log.Print("GetProductList was invoked")
	res, err := s.rSer.GetProductList()
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "repo.GetProductList went down witn error, cannot get products from db: %v\n", err)
	}

	resPb := convertProducts(res)

	response := &pb.GetProductListResponse{
		Result: resPb,
	}

	return response, nil
}

func enumSelect(i int) pb.ProductType {
	switch i {
	case 1:
		return pb.ProductType_PRODUCT_TYPE_SALAD
	case 2:
		return pb.ProductType_PRODUCT_TYPE_GARNISH
	case 3:
		return pb.ProductType_PRODUCT_TYPE_MEAT
	case 4:
		return pb.ProductType_PRODUCT_TYPE_SOUP
	case 5:
		return pb.ProductType_PRODUCT_TYPE_DRINK
	case 6:
		return pb.ProductType_PRODUCT_TYPE_DESSERT
	default:
		return pb.ProductType_PRODUCT_TYPE_UNSPECIFIED
	}

}

func convertProducts(res []*types.Product) []*pb.Product {
	var resPb []*pb.Product

	for _, v := range res {
		id := strconv.Itoa(v.Uuid)
		t := timestamppb.New(v.CreatedAt)
		pr := &pb.Product{
			Uuid:        id,
			Name:        v.Name,
			Description: v.Descript,
			Type:        enumSelect(v.Type),
			Weight:      int32(v.Weight),
			Price:       v.Price,
			CreatedAt:   t,
		}
		resPb = append(resPb, pr)
	}
	return resPb
}
