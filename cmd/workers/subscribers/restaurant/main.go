package main

import (
	app "github.com/MikhailMishutkin/FoodOrdering/cmd"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"log"
)

//type RestaurantSubscriber struct {
//	con *nats.Conn
//	natsrestaurant.NatsSub
//	natsrestservice.RestNATSService
//	repository.RestaurantRepo
//}
//
//func NewRS() *RestaurantSubscriber {
//	conn, err := nats.Connect(nats.DefaultURL)
//	if err != nil {
//		log.Fatal(err)
//	}
//	h := natsrestaurant.NewNATS()
//	return &RestaurantSubscriber{
//		con: conn,
//	}
//}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {
	conf, err := configs.New("./configs/main.yaml")
	if err != nil {
		log.Fatalf("can't receive config data: %v\n", err)
	}

	if err = app.StartRestSubs(conf); err != nil {
		log.Fatal(err)
	}
}

//init restaurant
//db, err := bootstrap.NewDB()
//if err != nil {
//	log.Fatal(err)
//}
//repo := repository.NewRestaurantRepo(db)

//order := &types.OrderRequest{}
//
//restSub := NewRS()
//
//sub, err := restSub.con.SubscribeSync("order")
//if err != nil {
//	fmt.Errorf("subscribeSync error: %v\n: ", err)
//}
//
//for {
//	//t := repository.DateConv(time.Now())
//	//t1 := t.AddDate(0, 0, 1)
//	//t2 := t1.Add(11 * time.Hour)
//	//t3 := t1.Add(21 * time.Hour)
//
//	msg, err := sub.NextMsgWithContext(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	if msg.Subject == "order" {
//		err = json.Unmarshal(msg.Data, order)
//		err := restSub.OrderReceive()
//		if err != nil {
//			return
//		}
//	}
//if time.Now().UnixNano() >= t2.UnixNano() && time.Now().UnixNano() < t3.UnixNano() {
//
//} else {
//	continue
//}
//}
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
