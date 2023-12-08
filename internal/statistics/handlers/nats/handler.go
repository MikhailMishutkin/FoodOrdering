package natsstatistics

import (
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type NatsSubStat struct {
	Conn          *nats.Conn
	Jm            NATSManager
	ClientProduct restaurant.ProductServiceClient
	ClientMenu    restaurant.MenuServiceClient
}

func NewNATSubStat(jm NATSManager, conf configs.Config) *NatsSubStat {
	nc, err := nats.Connect(nats.DefaultURL) //"nats:4222")
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}

	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	return &NatsSubStat{
		Conn:          nc,
		ClientProduct: restaurant.NewProductServiceClient(conn),
		ClientMenu:    restaurant.NewMenuServiceClient(conn),
		Jm:            jm,
	}
}

type NATSManager interface {
	DataSaveService(order *types.OrderRequest, products []*types.Product) error
}
