package app

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	handlerscustomer "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/grpc"
	cusrepository "github.com/MikhailMishutkin/FoodOrdering/internal/customer/repository"
	"github.com/MikhailMishutkin/FoodOrdering/internal/customer/service"
	handlers "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/handlers/grpc"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	serviceR "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/service"
	stathandlers "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/handlers/grpc"
	statrepository "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/repository"
	statservice "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/service"
	cust "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	rest "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/statistics"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"strings"
	//"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func StartGRPCAndHTTPServer(conf configs.Config) error {
	//connections to databases: pgx, gorm, sqlx
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	gorm, err := cusrepository.NewGormDB()
	if err != nil {
		log.Fatal("cannot connect to gorm: ", err)
	}

	dbx, err := statrepository.NewDB()
	if err != nil {
		log.Fatal("cannot connect to db on sqlx: ", err)
	}

	//connection to grpc
	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	//restaurant
	repo := repository.NewRestaurantRepo(db)
	ru := serviceR.NewRestaurantUsecace(repo)
	rs := handlers.NewRestaurantService(ru, cust.NewOfficeServiceClient(conn), cust.NewUserServiceClient(conn))

	//customer
	repoC := cusrepository.NewCustomerRepo(gorm)
	cu := service.NewCustomerUsecase(repoC)
	cs := handlerscustomer.New(rest.NewMenuServiceClient(conn), cu)

	//statistic
	repoS := statrepository.NewStatRepo(dbx)
	su := statservice.NewStatUsecase(repoS)
	sh := stathandlers.NewStatService(rest.NewProductServiceClient(conn), su)

	s := grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rest.RegisterProductServiceServer(s, &handlers.RestaurantService{})
	rest.RegisterMenuServiceServer(s, &handlers.RestaurantService{})
	rest.RegisterOrderServiceServer(s, &handlers.RestaurantService{})
	cust.RegisterOrderServiceServer(s, &handlerscustomer.CustomerService{})
	cust.RegisterOfficeServiceServer(s, &handlerscustomer.CustomerService{})
	cust.RegisterUserServiceServer(s, &handlerscustomer.CustomerService{})
	statistics.RegisterStatisticsServiceServer(s, &stathandlers.StatisticService{})

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

	err = statistics.RegisterStatisticsServiceHandlerServer(ctx, router, sh)
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
	//connections to databases: pgx, gorm, sqlx
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	gorm, err := cusrepository.NewGormDB()
	if err != nil {
		log.Fatal("cannot connect to gorm: ", err)
	}

	dbx, err := statrepository.NewDB()
	if err != nil {
		log.Fatal("cannot connect to db on sqlx: ", err)
	}

	//connection to grpc
	conn, err := grpc.Dial(conf.API.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	//restaurant
	repo := repository.NewRestaurantRepo(db)
	ru := serviceR.NewRestaurantUsecace(repo)
	rs := handlers.NewRestaurantService(ru, cust.NewOfficeServiceClient(conn), cust.NewUserServiceClient(conn))

	//customer
	repoC := cusrepository.NewCustomerRepo(gorm)
	cu := service.NewCustomerUsecase(repoC)
	n := handlerscustomer.New(rest.NewMenuServiceClient(conn), cu)

	//statistic
	repoS := statrepository.NewStatRepo(dbx)
	su := statservice.NewStatUsecase(repoS)
	sh := stathandlers.NewStatService(rest.NewProductServiceClient(conn), su)

	s := grpc.NewServer()

	rest.RegisterMenuServiceServer(s, rs)
	rest.RegisterProductServiceServer(s, rs)
	cust.RegisterOrderServiceServer(s, n)
	cust.RegisterOfficeServiceServer(s, n) // n - &handlers_customer.CustomerService{}
	cust.RegisterUserServiceServer(s, n)   // n - &handlers_customer.CustomerService{}
	statistics.RegisterStatisticsServiceServer(s, sh)
	lis, err := net.Listen("tcp", conf.API.GHost)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", conf.API.GHost)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
