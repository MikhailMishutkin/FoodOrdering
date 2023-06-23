package natscustomer

import (
	"fmt"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"log"
)

type NatsPub struct {
	js nats.JetStream
}

func NewNATS(js nats.JetStream) *NatsPub {
	return &NatsPub{
		js: js,
	}
}
func (np *NatsPub) NatsPublisher(order *pb.CreateOrderRequest) error {
	log.Println("customer publisher")
	fmt.Println(order)
	data, err := proto.Marshal(order)
	if err != nil {
		fmt.Errorf("cannot marshal proto message to binary: %w", err)
		return err
	}
	fmt.Println("after marshaling: ", data)

	return err
}

//_, err = js.AddConsumer("ORDERS", &nats.ConsumerConfig{
//Durable:      "my-consumer-1",
//Description:  "this is my awesome consumer",
//ReplayPolicy: nats.ReplayInstantPolicy,
//})
//fatalOnErr(err)
//
//sub, err := js.PullSubscribe("orders.us", "my-consumer-1")
//fatalOnErr(err)
//go processMsg(sub)
//
//time.Sleep(10 * time.Second)
//sub.Unsubscribe()
//
//log.Println("shutting down application...")
//}

//for {
//	resp, err := nc.Request("order", data, 500*time.Millisecond)
//	time.Sleep(1 * time.Second)Платежное поручение 00БП-000036 от 21.02.2023 16:22:57
//	if err != nil {
//		log.Printf("error sending message %v\n", err)
//		continue
//		log.Printf("reply: %v\n", string(resp.Data))
//	}
//
//}

//nc, err := nats.Connect(nats.DefaultURL)
//if err != nil {
//	log.Printf("can't connect to NATS: %v", err)
//	return err
//}
//defer nc.Drain()

//js, err := nc.JetStream()
//if err != nil {
//	log.Println("error line 24 natspublisher: ", err)
//}
//streamName := "ORDER"
//js.AddStream(&nats.StreamConfig{
//	Name:     streamName,
//	Subjects: []string{"orders.>"},
//})
