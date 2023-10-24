package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/internal/bootstrap"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	stathandlers "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/handlers/grpc"
	statrepository "github.com/MikhailMishutkin/FoodOrdering/internal/statistics/repository"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	log.Println("Statistics subscriber started")

	conf, err := configs.New("./configs/main.yaml.template")
	if err != nil {
		log.Fatal(err)
	}

	//init statistic
	dbx, err := bootstrap.NewDBX()
	if err != nil {
		log.Fatal("connect sqlx: ", err)
	}

	repoS := statrepository.NewStatRepo(dbx, conf)

	order := &types.OrderRequest{}

	sub, err := repoS.Conn.SubscribeSync("order")
	if err != nil {
		fmt.Errorf("subscribeSync error: %v\n: ", err)
	}

	for {
		q, err := sub.Delivered()
		if err != nil {
			log.Println(err)
		}
		if q > 0 {
			t := repository.DateConv(time.Now())
			t1 := t.AddDate(0, 0, 1)
			t2 := timestamppb.New(t1)
			gMenuReq := restaurant.GetMenuRequest{
				OnDate: t2,
			}
			menuResp, err := repoS.ClientMenu.GetMenu(context.Background(), &gMenuReq)
			if err != nil {
				fmt.Errorf("restaurant GetMenu error: %v\n", err)
			}
			start := stathandlers.TimeAssert(menuResp.Menu.OpeningRecordAt)
			end := stathandlers.TimeAssert(menuResp.Menu.ClosingRecordAt)

			if time.Now().UnixNano() >= start.UnixNano() && time.Now().UnixNano() < end.UnixNano() {
				msg, err := sub.NextMsgWithContext(context.Background())
				if err != nil {
					log.Fatal(err)
				}
				if msg.Subject == "order" {
					err = json.Unmarshal(msg.Data, order)
					repoS.ReceiveOrdersRepo(order)
				}
			} else {
				continue
			}
		} else {
			continue
		}

	}
}

//su := statservice.NewStatUsecase(repoS)
//ss := stathandlers.NewStatService(, su)

//nc, err := nats.Connect(nats.DefaultURL)
//if err != nil {
//	log.Fatalf("can't connect to NATS: %v", err)
//}
//defer nc.Close()

//connection to grpc
//conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
//if err != nil {
//	log.Fatalf("Failed to connect: %v\n", err)
//}
//defer conn.Close()
