package handlers

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
	"github.com/golang/protobuf/proto"
	"log"

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
		//p = &pb.Product{
		//	//Uuid:        gen.RandomID(),
		//	Name:        in.Name,
		//	Description: in.Description,
		//	Type:        in.Type,
		//	Weight:      in.Weight,
		//	Price:       in.Price,
		//	CreatedAt:   timestamppb.Now(),
		//}
		prData, err := proto.Marshal(in)
		if err != nil {
			code := codes.Internal
			return nil, status.Errorf(code, "repo.CreateProduct can't marshal data: %v\n", err)
		}

		err = s.rSer.CreateProduct(prData)
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
	response := &pb.GetProductListResponse{}
	err = proto.Unmarshal(res, response)
	if err != nil {
		code := codes.DataLoss
		return nil, status.Errorf(code, "can't unmarshal protomessage with GetProductListResponse: %v\n", err)
	}

	return response, nil
}
