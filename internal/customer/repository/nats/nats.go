package natscustomerrepo

import (
	"github.com/nats-io/nats.go"
	"log"
)

type NatsPublisherRepo struct {
	conn *nats.Conn
}

func NewNPublisherRepo() *NatsPublisherRepo {
	//conf, err := configs.New("./configs/main.yaml.template")
	//if err != nil {
	//	fmt.Errorf("Can't load config in publisher repo: %v\n", err)
	//}
	nc, err := nats.Connect(nats.DefaultURL) //"nats:4222")
	if err != nil {
		log.Println("can't connect to NATS-server:", err)
	}
	return &NatsPublisherRepo{
		conn: nc,
	}
}
