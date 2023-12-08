package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/internal/bootstrap"
	handlerscustomer "github.com/MikhailMishutkin/FoodOrdering/internal/customer/handlers/grpc"
	natscustomerrepo "github.com/MikhailMishutkin/FoodOrdering/internal/customer/repository/nats"
	cusrepository "github.com/MikhailMishutkin/FoodOrdering/internal/customer/repository/postgres"
	service "github.com/MikhailMishutkin/FoodOrdering/internal/customer/service/grpc"
	handlers "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/handlers/grpc"
	natsrestaurant "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/handlers/nats"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository/postgres"
	serviceR "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/service/grpc"
	natsrestservice "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/service/nats"
	stathandlers "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/handlers/grpc"
	natsstatistics "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/handlers/nats"
	statrepository "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/repository/postgres"
	statservice "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/service/grpc"
	natsstatservice "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/service/nats"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
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
		return fmt.Errorf("cannot connect to db on pqx: %v\n ", err)
	}

	gorm, err := bootstrap.NewGormDB()
	if err != nil {
		return fmt.Errorf("cannot connect to gorm: %v\n ", err)
	}

	dbx, err := bootstrap.NewDBX()
	if err != nil {
		return fmt.Errorf("cannot connect to db on sqlx: %v\n", err)
	}

	//connection to grpc
	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Failed to create gRPC client connection: %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	//restaurant
	repo := repository.NewRestaurantRepo(db)
	ru := serviceR.NewRestaurantUsecase(repo)
	rs := handlers.NewRestaurantService(ru)

	//customer
	repoC := cusrepository.NewCustomerRepo(gorm)
	natsC := natscustomerrepo.NewNPublisherRepo()
	cu := service.NewCustomerUsecase(repoC, natsC)
	cs := handlerscustomer.New(restaurant.NewMenuServiceClient(conn), cu)

	//statistic
	repoS := statrepository.NewStatRepo(dbx, conf)
	su := statservice.NewStatUsecase(repoS)
	sh := stathandlers.NewStatService(su)

	s := grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	restaurant.RegisterProductServiceServer(s, &handlers.RestaurantService{})
	restaurant.RegisterMenuServiceServer(s, &handlers.RestaurantService{})
	restaurant.RegisterOrderServiceServer(s, &handlers.RestaurantService{})
	customer.RegisterOrderServiceServer(s, cs)
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

	err = customer.RegisterOrderServiceHandlerServer(ctx, router, cs)
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
		return fmt.Errorf("cannot connect to db on pqx: %v\n ", err)
	}

	gorm, err := bootstrap.NewGormDB()
	if err != nil {
		return fmt.Errorf("cannot connect to gorm: %v\n ", err)
	}

	dbx, err := bootstrap.NewDBX()
	if err != nil {
		return fmt.Errorf("cannot connect to db on sqlx: %v\n", err)
	}

	//connection to grpc
	conn, err := grpc.Dial(conf.API.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Failed to create gRPC client connection: %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	//restaurant
	repo := repository.NewRestaurantRepo(db)
	ru := serviceR.NewRestaurantUsecase(repo)
	rs := handlers.NewRestaurantService(ru)

	//customer
	repoC := cusrepository.NewCustomerRepo(gorm)
	natsC := natscustomerrepo.NewNPublisherRepo()
	cu := service.NewCustomerUsecase(repoC, natsC)
	n := handlerscustomer.New(restaurant.NewMenuServiceClient(conn), cu)

	//statistic
	repoS := statrepository.NewStatRepo(dbx, conf)
	su := statservice.NewStatUsecase(repoS)
	sh := stathandlers.NewStatService(su)

	s := grpc.NewServer()

	restaurant.RegisterMenuServiceServer(s, rs)
	restaurant.RegisterProductServiceServer(s, rs)
	restaurant.RegisterOrderServiceServer(s, rs)
	customer.RegisterOrderServiceServer(s, n)
	customer.RegisterOfficeServiceServer(s, n) // n - &handlers_customer.CustomerService{}
	customer.RegisterUserServiceServer(s, n)   // n - &handlers_customer.CustomerService{}
	statistics2.RegisterStatisticsServiceServer(s, sh)

	router := runtime.NewServeMux()
	err = restaurant.RegisterOrderServiceHandlerServer(context.Background(), router, rs)
	if err != nil {
		log.Printf("Failed to register gateway: %v\n", err)
	}

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

func StartRestSubs(conf configs.Config) error {
	log.Println("Restaurant subscriber was started")
	db, err := bootstrap.NewDB()
	if err != nil {
		return fmt.Errorf("cannot connect to db on pqx: %v\n", err)
	}

	//connection to grpc
	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Failed to create gRPC client connection: %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	//natsrestaurant
	repo1 := repository.NewRestaurantRepo(db)
	serv1 := natsrestservice.NewRNService(repo1)
	handl := natsrestaurant.NewNATS(serv1, customer.NewOfficeServiceClient(conn), customer.NewUserServiceClient(conn))

	order := &types.OrderRequest{}

	sub, err := handl.Conn.SubscribeSync("order")
	if err != nil {
		return fmt.Errorf("subscribeSync error: %v\n: ", err)
	}

	for {
		//t := repository.DateConv(time.Now())
		//t1 := t.AddDate(0, 0, 1)
		//t2 := t1.Add(11 * time.Hour)
		//t3 := t1.Add(21 * time.Hour)

		msg, err := sub.NextMsgWithContext(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if msg.Subject == "order" {
			err = json.Unmarshal(msg.Data, order)
			err := handl.OrderReceive(order)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func StartStatSubs(conf configs.Config) error {
	log.Println("Statistics subscriber was started")
	dbx, err := bootstrap.NewDBX()
	if err != nil {
		return fmt.Errorf("cannot connect to db on sqlx: %v\n", err)
	}

	//connection to grpc
	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("Failed to create gRPC client connection: %v\n", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	//natsstatistic
	repo1 := statrepository.NewStatRepo(dbx, conf)
	serv1 := natsstatservice.NewSNService(repo1)
	handl := natsstatistics.NewNATSubStat(serv1, conf)

	order := &types.OrderRequest{}

	sub, err := handl.Conn.SubscribeSync("order")
	if err != nil {
		return fmt.Errorf("subscribeSync error: %v\n: ", err)
	}

	for {
		//q, err := sub.Delivered()
		//		if err != nil {
		//			log.Println(err)
		//		}
		//		if q > 0 {
		//		//t := repository.DateConv(time.Now())
		//		//t1 := t.AddDate(0, 0, 1)
		//		//start := t1.Add(11 * time.Hour)
		//		//end := t1.Add(21 * time.Hour)
		//
		//if time.Now().UnixNano() >= start.UnixNano() && time.Now().UnixNano() < end.UnixNano() {

		msg, err := sub.NextMsgWithContext(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if msg.Subject == "order" {
			err = json.Unmarshal(msg.Data, order)
			err := handl.OrderReceive(order)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
