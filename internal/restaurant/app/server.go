package app

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	"github.com/MikhailMishutkin/FoodOrdering/proto"

	//"github.com/goccy/go-json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc"
)

//var collection *mongo.Collection
var addr string = "0.0.0.0:8080"

type Server struct {
	pb.UnimplementedProductServiceServer
	pb.UnimplementedMenuServiceServer
	pb.UnimplementedOrderServiceServer
	repo RestaurantRepository
}

type RestaurantRepository interface {
	CreateProduct(p *pb.Product) error
	GetProductList() (*pb.GetProductListResponse, error)
	CreateMenu() (*pb.Menu, error)
	GetMenu(time.Time) *pb.Menu
	GetOrderList() ([]*pb.Order, []*pb.OrdersByOffice)
}

func NewServer(rp RestaurantRepository) *Server {
	return &Server{repo: rp}
}

func StartGRPCRestaurantServer() error {
	repo := repository.NewRestaurantRepo()
	uc := NewServer(repo)
	s := grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pb.RegisterProductServiceServer(s, &Server{})
	pb.RegisterMenuServiceServer(s, &Server{})
	pb.RegisterOrderServiceServer(s, &Server{})

	router := runtime.NewServeMux()
	err := proto.RegisterProductServiceHandlerServer(ctx, router, uc)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}
	err = proto.RegisterMenuServiceHandlerServer(ctx, router, uc)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	log.Printf("Starting GRPCRestaurantServer at port: %v\n ", addr)
	return http.ListenAndServe(addr, httpGrpcRouter(s, router))

}

func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}

//TODO: перенести в отдельный файл Order
func (s *Server) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	log.Print("GetUpToDateOrderList was invoked")
	s.repo.GetOrderList()
	return nil, nil
}
