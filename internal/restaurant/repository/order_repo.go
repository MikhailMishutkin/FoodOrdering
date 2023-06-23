package repository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	pbr "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

//TODO: ошибки!!!!
func (r *RestaurantRepo) GetOrderList() ([]*pbr.Order, []*pbr.OrdersByOffice) {
	log.Println("restaurant subscriber")
	config, err := configs.New("./configs/main.yaml.template")
	if err != nil {
		log.Println("can't load config to restaurant subscriber: ", err)
	}
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Println(err)
	}

	_, err = js.AddConsumer(config.NATS.Name, &nats.ConsumerConfig{
		Durable: config.NATS.Consumer,
		//DeliverSubject: "orders",
		AckPolicy: nats.AckExplicitPolicy,
		AckWait:   time.Second,
	})
	if err != nil {
		log.Println("can't add consumer restaurant: ", err)
	}

	sub, err := js.PullSubscribe("orders", config.NATS.Consumer) // nats.Bind(config.NATS.Name, config.NATS.Consumer) если Sync
	if err != nil {
		log.Println(err)
	}

	msgs, err := sub.Fetch(5)
	if err != nil {
		log.Printf("msg queue can't read: %v\n", err)
	}

	m := make(map[string]int64)
	var so []*pbr.Order

	for _, msg := range msgs {
		var co pb.CreateOrderRequest

		err = proto.Unmarshal(msg.Data, &co)
		if err != nil {
			log.Println(err)
		}
		m = fieldsFromOrderItem(co.Salads, m)

		m = fieldsFromOrderItem(co.Garnishes, m)

		m = fieldsFromOrderItem(co.Meats, m)

		m = fieldsFromOrderItem(co.Soups, m)

		m = fieldsFromOrderItem(co.Drinks, m)

		m = fieldsFromOrderItem(co.Desserts, m)

	}
	for k, v := range m {
		fmt.Printf("productId: %v, count: %v\n", k, v)
		o := &pbr.Order{
			ProductId: k,
			Count:     v,
		}
		so = append(so, o)
		fmt.Println(so)
	}

	//sub.Unsubscribe()

	return so, nil
}

func fieldsFromOrderItem(oi []*pb.OrderItem, m map[string]int64) map[string]int64 {
	o := &pbr.Order{}

	for _, v := range oi {
		if v.Count != 0 {
			o.ProductId = v.ProductUuid
			o.Count = int64(v.Count)
			if check(m, v.ProductUuid) {
				a := m[o.ProductId]
				o.Count = a + int64(v.Count)
				m[o.ProductId] = o.Count

				continue
			}

			m[o.ProductId] = int64(v.Count)

		} else {
			continue
		}
	}

	return m
}

func check(m map[string]int64, id string) bool {
	for k, _ := range m {
		if k == id {
			return true
		} else {
			continue
		}
	}
	return false
}

//msg, err := sub.NextMsg(time.Second)
//if err != nil {
//	log.Println(err)
//}
//fmt.Printf("received %q\n", msg.Subject, msg.Data)

//err = msg.Ack()
//if err != nil {
//	log.Printf("can't send Ack: %v\n", err)
//}
//queuedMsgs, _, err := sub.Pending()
//if err != nil {
//	log.Println(err)
//}
//log.Printf("%d messages queued\n", queuedMsgs)
