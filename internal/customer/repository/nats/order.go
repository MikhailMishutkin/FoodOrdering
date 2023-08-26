package natscustomerrepo

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/nats-io/nats.go"
	"log"
)

func (np *NatsPublisherRepo) NatsPublisher(order *types.OrderRequest) error {
	log.Println("NatsPublisher repo was invoked")

	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("Can't marshal to bytes: %v\n", err)
	}

	msg := &nats.Msg{
		Subject: "order",
		Data:    data,
	}
	np.conn.PublishMsg(msg)

	return nil
}
