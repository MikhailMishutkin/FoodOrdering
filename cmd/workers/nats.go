package main

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	fmt.Println("restaurant subscriber started")
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
			repo.GetOrder(order)
		} else {
			continue
		}
	}
}
