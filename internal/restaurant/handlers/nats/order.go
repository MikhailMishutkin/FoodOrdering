package natsrestaurant

//
//import (
//	"fmt"
//	"github.com/MikhailMishutkin/FoodOrdering/configs"
//	serviceR "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/service"
//	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
//	"github.com/golang/protobuf/proto"
//	"github.com/nats-io/nats.go"
//	"log"
//	"time"
//)
//
//type NatsSub struct {
//	js nats.JetStream
//}
//
//func NewNATS(js nats.JetStream) *NatsSub {
//	return &NatsSub{
//		js: js,
//	}
//}
//
////TODO: обработать ошибки!
//func (ns *NatsSub) NatsSubscriber() error {
//	log.Println("restaurant subscriber")
//	config, err := configs.New("./config")
//	if err != nil {
//		return err
//	}
//
//	sub, err := ns.js.SubscribeSync("", nats.Bind(config.NATS.Name, config.NATS.Consumer))
//	if err != nil {
//		log.Println("line 33 ordernats: ", err)
//	}
//	msg, err := sub.NextMsg(time.Second)
//	if err != nil {
//		log.Println("line 37 ordernats: ", err)
//	}
//	fmt.Printf("received %q\n", msg.Subject)
//
//	msg.Ack()
//	queuedMsgs, _, err := sub.Pending()
//	if err != nil {
//		log.Println("line 44 ordernats: ", err)
//	}
//	fmt.Printf("%d messages queued\n", queuedMsgs)
//	var co *pb.CreateOrderRequest
//	err = proto.Unmarshal(msg.Data, co)
//	if err != nil {
//		log.Println("line 50 ordernats: ", err)
//	}
//	fmt.Println("rest", co)
//	(&serviceR.RestaurantUsecase{}).GetOrderList()
//	//sub.Unsubscribe()
//
//	return err
//
//}

//streamName := "ORDER"
//js.AddStream(&nats.StreamConfig{
//Name:     streamName,
//Subjects: []string{"orders.>"},
//})
//
//js.Subscribe("orders.>", func(m *nats.Msg) {
//	fmt.Print(m.Data)
//}, nats.BindStream("ORDER"))
//
//consumerName := "restaurant"
//js.AddConsumer(conf.NATS.Name, &nats.ConsumerConfig{
//Durable: consumerName,
////DeliverSubject: "handler-2",
//AckPolicy: nats.AckExplicitPolicy,
//AckWait:   time.Second,
//})
//
//_ = natscustomer.NewNATS(js)
//
//
//return err

//nc, err := nats.Connect(nats.DefaultURL)
//if err != nil {
//	log.Fatalf("can't connect to NATS: %v", err)
//}
//defer nc.Close()
//
//js, err := nc.JetStream()
//if err != nil {
//	log.Println("line 20 ordernats: ", err)
//}
