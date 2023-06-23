package app

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	handlers_customer "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/grpc"
	cusrepository "github.com/MikhailMishutkin/FoodOrdering/internal/customer/repository"
	"github.com/MikhailMishutkin/FoodOrdering/internal/customer/service"
	handlers "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/handlers/grpc"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	serviceR "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/service"
	cust "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	rest "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"strings"
	//"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func StartGRPCAndHTTPServer(conf configs.Config) error {
	repo := repository.NewRestaurantRepo()
	ru := serviceR.NewRestaurantUsecace(repo)
	rs := handlers.NewRestaurantService(ru)

	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	repoC := cusrepository.NewCustomerRepo()
	cu := service.NewCustomerUsecase(repoC)
	cs := handlers_customer.New(rest.NewMenuServiceClient(conn), cu)

	s := grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rest.RegisterProductServiceServer(s, &handlers.RestaurantService{})
	rest.RegisterMenuServiceServer(s, &handlers.RestaurantService{})
	rest.RegisterOrderServiceServer(s, &handlers.RestaurantService{})
	cust.RegisterOrderServiceServer(s, &handlers_customer.CustomerService{})
	cust.RegisterOfficeServiceServer(s, &handlers_customer.CustomerService{})
	cust.RegisterUserServiceServer(s, &handlers_customer.CustomerService{})

	router := runtime.NewServeMux()
	err = rest.RegisterProductServiceHandlerServer(ctx, router, rs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}
	err = rest.RegisterMenuServiceHandlerServer(ctx, router, rs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = rest.RegisterOrderServiceHandlerServer(ctx, router, rs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = cust.RegisterOrderServiceHandlerServer(ctx, router, cs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = cust.RegisterOfficeServiceHandlerServer(ctx, router, cs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = cust.RegisterUserServiceHandlerServer(ctx, router, cs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	log.Printf("Starting GRPCRestaurantServer at port: %v\n ", conf.API.Host)
	return http.ListenAndServe(conf.API.Host, httpGrpcRouter(s, router))

}

// TODO: проверить как работает без этого роутера, может из-за него не получается выполнить запрос в рамках одного сервера
func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}

func StartGRPC(conf configs.Config) {
	repo := repository.NewRestaurantRepo()
	ru := serviceR.NewRestaurantUsecace(repo)
	rs := handlers.NewRestaurantService(ru)
	conn, err := grpc.Dial(conf.API.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	repoC := cusrepository.NewCustomerRepo()
	cu := service.NewCustomerUsecase(repoC)
	n := handlers_customer.New(rest.NewMenuServiceClient(conn), cu)

	s := grpc.NewServer()

	rest.RegisterMenuServiceServer(s, rs)

	cust.RegisterOrderServiceServer(s, n)

	lis, err := net.Listen("tcp", conf.API.GHost)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", conf.API.GHost)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
