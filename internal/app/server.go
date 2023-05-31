package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"

	hc "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/grpc"
	cusrepository "github.com/MikhailMishutkin/FoodOrdering/internal/customer/repository"
	handlers "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/handlers/grpc"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	cust "github.com/MikhailMishutkin/FoodOrdering/pkg/contracts-v0.3.0/pkg/contracts/customer"
	rest "github.com/MikhailMishutkin/FoodOrdering/pkg/contracts-v0.3.0/pkg/contracts/restaurant"

	//"github.com/goccy/go-json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:8080"
var gAddr string = "0.0.0.0:5051"

func StartGRPCAndHTTPServer() error {
	repo := repository.NewRestaurantRepo()
	rs := handlers.NewRestaurantService(repo)

	repoC := cusrepository.NewCustomerRepo()
	cs := hc.NewCustomerService(repoC)

	opts := []grpc.ServerOption{}
	// tls := true //change that to false if needed

	// if tls {
	// 	certFile := "ssl/server.crt"
	// 	keyFile := "ssl/server.pem"
	// 	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

	// 	if err != nil {
	// 		log.Fatalf("Failed loading certificates: %v\n", err)
	// 	}

	// 	opts = append(opts, grpc.Creds(creds))
	// }

	s := grpc.NewServer(opts...)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rest.RegisterProductServiceServer(s, rest.UnimplementedProductServiceServer{})
	rest.RegisterMenuServiceServer(s, rest.UnimplementedMenuServiceServer{})
	rest.RegisterOrderServiceServer(s, rest.UnimplementedOrderServiceServer{})
	cust.RegisterOrderServiceServer(s, cust.UnimplementedOrderServiceServer{})

	router := runtime.NewServeMux()
	err := rest.RegisterProductServiceHandlerServer(ctx, router, rs)
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
	// lis, err := net.Listen("tcp", addr)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s.Serve(lis)
	log.Printf("Starting GRPCRestaurantServer at port: %v\n ", addr)
	return http.ListenAndServe(addr, httpGrpcRouter(s, router))

}

func StartGRPC() {

	lis, err := net.Listen("tcp", gAddr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", gAddr)

	s := grpc.NewServer()

	rest.RegisterMenuServiceServer(s, rest.UnimplementedMenuServiceServer{})
	//reflection.Register(s) //подтянули рефлексию для использования Evans после запуска сервера evans --host localhost --port 50051 --reflection repl

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
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
