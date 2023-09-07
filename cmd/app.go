package app

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/internal/bootstrap"
	handlerscustomer "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/grpc"
	natscustomer "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/nats"
	natscustomerrepo "github.com/MikhailMishutkin/FoodOrdering/internal/customer/repository/nats"
	cusrepository "github.com/MikhailMishutkin/FoodOrdering/internal/customer/repository/postgres"
	service "github.com/MikhailMishutkin/FoodOrdering/internal/customer/service/grpc"
	natscustomerservice "github.com/MikhailMishutkin/FoodOrdering/internal/customer/service/nats"
	handlers "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/handlers/grpc"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	serviceR "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/service/grpc"
	stathandlers "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/handlers/grpc"
	statrepository "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/repository"
	statservice "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/service"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	statistics2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/statistics"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"
)

func StartGRPCAndHTTPServer(conf configs.Config) error {
	//connections to databases: pgx, gorm, sqlx
	db, err := bootstrap.NewDB()
	if err != nil {
		return fmt.Errorf("cannot connect to db on pqx: ", err)
	}

	gorm, err := bootstrap.NewGormDB()
	if err != nil {
		return fmt.Errorf("cannot connect to gorm: ", err)
	}

	dbx, err := bootstrap.NewDBX()
	if err != nil {
		return fmt.Errorf("cannot connect to db on sqlx: ", err)
	}

	//connection to grpc
	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Failed to create gRPC client connection: %v\n", err)
	}
	defer conn.Close()

	//restaurant
	repo := repository.NewRestaurantRepo(db, conf)
	ru := serviceR.NewRestaurantUsecace(repo)
	rs := handlers.NewRestaurantService(ru, customer.NewOfficeServiceClient(conn), customer.NewUserServiceClient(conn))

	//customer
	repoC := cusrepository.NewCustomerRepo(gorm)
	cu := service.NewCustomerUsecase(repoC)
	cs := handlerscustomer.New(restaurant.NewMenuServiceClient(conn), cu)

	//statistic
	repoS := statrepository.NewStatRepo(dbx, conf)
	su := statservice.NewStatUsecase(repoS)
	sh := stathandlers.NewStatService(su)

	//natscustomer
	natsP := natscustomerrepo.NewNPublisherRepo()
	natsU := natscustomerservice.NewCNU(natsP)
	natsH := natscustomer.NewNATS(natsU)

	s := grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	restaurant.RegisterProductServiceServer(s, &handlers.RestaurantService{})
	restaurant.RegisterMenuServiceServer(s, &handlers.RestaurantService{})
	restaurant.RegisterOrderServiceServer(s, &handlers.RestaurantService{})
	customer.RegisterOrderServiceServer(s, natsH)
	customer.RegisterOfficeServiceServer(s, &handlerscustomer.CustomerService{})
	customer.RegisterUserServiceServer(s, &handlerscustomer.CustomerService{})
	statistics2.RegisterStatisticsServiceServer(s, &stathandlers.StatisticService{})

	router := runtime.NewServeMux()

	err = restaurant.RegisterProductServiceHandlerServer(ctx, router, rs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}
	err = restaurant.RegisterMenuServiceHandlerServer(ctx, router, rs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = restaurant.RegisterOrderServiceHandlerServer(ctx, router, rs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = customer.RegisterOrderServiceHandlerServer(ctx, router, natsH)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = customer.RegisterOfficeServiceHandlerServer(ctx, router, cs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = customer.RegisterUserServiceHandlerServer(ctx, router, cs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

	err = statistics2.RegisterStatisticsServiceHandlerServer(ctx, router, sh)
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

func StartGRPC(conf configs.Config) error {
	//connections to databases: pgx, gorm, sqlx
	db, err := bootstrap.NewDB()
	if err != nil {
		return fmt.Errorf("cannot connect to db on pqx: ", err)
	}

	gorm, err := bootstrap.NewGormDB()
	if err != nil {
		return fmt.Errorf("cannot connect to gorm: ", err)
	}

	dbx, err := bootstrap.NewDBX()
	if err != nil {
		return fmt.Errorf("cannot connect to db on sqlx: ", err)
	}

	//connection to grpc
	conn, err := grpc.Dial(conf.API.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Failed to create gRPC client connection: %v\n", err)
	}
	defer conn.Close()

	//restaurant
	repo := repository.NewRestaurantRepo(db, conf)
	ru := serviceR.NewRestaurantUsecace(repo)
	rs := handlers.NewRestaurantService(ru, customer.NewOfficeServiceClient(conn), customer.NewUserServiceClient(conn))

	//customer
	repoC := cusrepository.NewCustomerRepo(gorm)
	cu := service.NewCustomerUsecase(repoC)
	n := handlerscustomer.New(restaurant.NewMenuServiceClient(conn), cu)

	//statistic
	repoS := statrepository.NewStatRepo(dbx, conf)
	su := statservice.NewStatUsecase(repoS)
	sh := stathandlers.NewStatService(su)

	s := grpc.NewServer()

	restaurant.RegisterMenuServiceServer(s, rs)
	restaurant.RegisterProductServiceServer(s, rs)
	//customer.RegisterOrderServiceServer(s, n)
	customer.RegisterOfficeServiceServer(s, n) // n - &handlers_customer.CustomerService{}
	customer.RegisterUserServiceServer(s, n)   // n - &handlers_customer.CustomerService{}
	statistics2.RegisterStatisticsServiceServer(s, sh)
	lis, err := net.Listen("tcp", conf.API.GHost)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", conf.API.GHost)

	if err = s.Serve(lis); err != nil {
		return fmt.Errorf("Failed to serve: %v\n", err)
	}
	return nil
}
