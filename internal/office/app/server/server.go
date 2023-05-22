package server_office

import (
	"context"
	"log"
	"net/http"
	"strings"

	cusrepository "github.com/MikhailMishutkin/FoodOrdering/internal/office/repository"
	"github.com/MikhailMishutkin/FoodOrdering/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:8081"

type ServerCustomer struct {
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedOfficeServiceServer

	repo CustomerRepository
}

type CustomerRepository interface {
	GetMenu()
}

func NewServer(cp CustomerRepository) *ServerCustomer {
	return &ServerCustomer{repo: cp}
}

func StartGRPCCustomerServer() error {
	repo := cusrepository.NewCustomerRepo()
	uc := NewServer(repo)
	s := grpc.NewServer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pb.RegisterOrderServiceServer(s, &ServerCustomer{})
	//  pb.RegisterMenuServiceServer(s, &Server{})
	// pb.RegisterOrderServiceServer(s, &Server{})

	router := runtime.NewServeMux()
	err := proto.RegisterOrderServiceHandlerServer(ctx, router, uc)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}
	// err = proto.RegisterMenuServiceHandlerServer(context.Background(), router, uc)
	// if err != nil {
	// 	log.Printf("Failed to register gateway: %v\n", err)
	// }

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

func (cs *ServerCustomer) GetActualMenu(ctx context.Context, in *pb.GetActualMenuRequest) (*pb.GetActualMenuResponse, error) {
	go cs.repo.GetMenu()
	a := &pb.GetActualMenuResponse{}
	return a, nil
}
