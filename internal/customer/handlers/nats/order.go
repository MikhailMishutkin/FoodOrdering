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

	data, err := proto.Marshal(order)
	if err != nil {
		fmt.Errorf("cannot marshal proto message to binary: %w", err)
		return err
	}

	np.js.Publish("orders.>", data)

	return err
}

//for {
//	resp, err := nc.Request("order", data, 500*time.Millisecond)
//	time.Sleep(1 * time.Second)
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
