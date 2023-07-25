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
