package cusrepository

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/nats-io/nats.go"
	"log"
)

func (cr *CustomerRepo) CreateOrder(order *types.OrderRequest) error {
	log.Println("CreateOrder repo was invoked")

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("can't connect to NATS: %v", err)
		return err
	}
	defer nc.Drain()

	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("Can't marshal to bytes: %v\n", err)
	}

	msg := &nats.Msg{
		Subject: "order",
		Data:    data,
	}
	nc.PublishMsg(msg)

	return nil
}
