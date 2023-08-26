package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/bootstrap"
	"log"
	"time"

	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
)

func main() {
	fmt.Println("Restaurant subscriber started")

	conf, err := configs.New("./configs/main.yaml.template")
	if err != nil {
		log.Fatal(err)
	}
	//init restaurant
	db, err := bootstrap.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewRestaurantRepo(db, conf)

	order := &types.OrderRequest{}

	sub, err := repo.Conn.SubscribeSync("order")
	if err != nil {
		fmt.Errorf("subscribeSync error: %v\n: ", err)
	}

	for {
		t := repository.DateConv(time.Now())
		t1 := t.AddDate(0, 0, 1)
		menu, err := repo.GetMenu(t1)
		if err != nil {
			fmt.Errorf("restaurant GetMenu error: %v\n", err)
		}
		if time.Now().UnixNano() >= menu.OpenAt.UnixNano() && time.Now().UnixNano() < menu.ClosedAt.UnixNano() {
			msg, err := sub.NextMsgWithContext(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			if msg.Subject == "order" {
				err = json.Unmarshal(msg.Data, order)
				repo.ReceiveOrder(order)
			}
		} else {
			continue
		}

	}

}

//nrs := natsrestaurantservice.NewNRU()
//nrh := natsrestaurant.NewRestSubs(nrs)

//conf, err := configs.New("./configs/main.yaml.template")
//if err != nil {
//	log.Fatal(err)
//}
//connection to grpc
//conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
//if err != nil {
//	log.Fatalf("Failed to connect: %v\n", err)
//}
//defer conn.Close()
