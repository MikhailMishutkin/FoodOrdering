package natsrestaurant

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

//TODO: обработать ошибки!
func NatsSubscriber() error {
	log.Println("restaurant subscriber")
	config, err := configs.New("./config")
	if err != nil {
		return err
	}
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Println("line 20 ordernats: ", err)
	}
	consumerName := "restaurant"
	js.AddConsumer(config.NATS.Name, &nats.ConsumerConfig{
		Durable: consumerName,
		//DeliverSubject: "handler-2",
		AckPolicy: nats.AckExplicitPolicy,
		AckWait:   time.Second,
	})

	sub, err := js.SubscribeSync("", nats.Bind(config.NATS.Name, consumerName))
	if err != nil {
		log.Println("line 33 ordernats: ", err)
	}
	msg, err := sub.NextMsg(time.Second)
	if err != nil {
		log.Println("line 37 ordernats: ", err)
	}
	fmt.Printf("received %q\n", msg.Subject)

	msg.Ack()
	queuedMsgs, _, err := sub.Pending()
	if err != nil {
		log.Println("line 44 ordernats: ", err)
	}
	fmt.Printf("%d messages queued\n", queuedMsgs)
	var co *pb.CreateOrderRequest
	err = proto.Unmarshal(msg.Data, co)
	if err != nil {
		log.Println("line 50 ordernats: ", err)
	}
	fmt.Println("rest", co)
	//sub.Unsubscribe()

	return err

}
