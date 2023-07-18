package handlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
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
