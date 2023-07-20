package main

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	stathandlers "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/handlers/grpc"
	statrepository "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/repository"
	statservice "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/service"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func init() {

}

func main() {
	fmt.Println("restaurant subscriber started")

	conf, err := configs.New("./configs/main.yaml.template")
	//init stat
	dbx, err := statrepository.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	//connection to grpc
	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	//statistic
	repoS := statrepository.NewStatRepo(dbx)
	su := statservice.NewStatUsecase(repoS)
	ss := stathandlers.NewStatService(restaurant.NewProductServiceClient(conn), su)

	//init rest
	db, err := repository.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	order := &types.OrderRequest{}
	repo := &repository.RestaurantRepo{
		DB: db,
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	sub, err := nc.SubscribeSync("order")
	if err != nil {
		fmt.Errorf("subscribeSync error: %v\n: ", err)
	}
	for {
		msg, _ := sub.NextMsg(10 * time.Minute)
		if msg.Subject == "order" {
			err = json.Unmarshal(msg.Data, order)
			(&stathandlers.StatisticService{
				ClientProduct: restaurant.NewProductServiceClient(conn),
				SS:            ss.SS,
			}).GetProduct()
			(&stathandlers.StatisticService{
				SS: ss.SS,
			}).GetOrders(order)
			repo.GetOrder(order)

		} else {
			continue
		}
	}
}
