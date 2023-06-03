package restaurant_app

import (
	"context"
	"log"
	"net/http"
	"strings"

	hc "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/grpc"
	cusrepository "github.com/MikhailMishutkin/FoodOrdering/internal/customer/repository"
	handlers "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/handlers/grpc"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	cust "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	rest "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"

	//"github.com/goccy/go-json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:8080"

func StartGRPCAndHTTPServer() error {
	repo := repository.NewRestaurantRepo()
	//log.Print("check repo: ", repo)
	rs := handlers.NewRestaurantService(repo)

	repoC := cusrepository.NewCustomerRepo()
	cs := hc.NewCustomerService(repoC)

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rest.RegisterProductServiceServer(s, &handlers.RestaurantService{})
	rest.RegisterMenuServiceServer(s, &handlers.RestaurantService{})
	rest.RegisterOrderServiceServer(s, &handlers.RestaurantService{})
	cust.RegisterOrderServiceServer(s, &hc.CustomerService{})

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

	log.Printf("Starting GRPCRestaurantServer at port: %v\n ", addr)
	return http.ListenAndServe(addr, httpGrpcRouter(s, router))

}

//TODO: проверить как работает без этого роутера
func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}

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

// lis, err := net.Listen("tcp", addr)
// if err != nil {
// 	log.Fatalf("failed to listen: %v", err)
// }
// s.Serve(lis)
