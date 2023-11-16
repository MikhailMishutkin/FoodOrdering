package natsrestaurant

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	customer2 "github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"github.com/nats-io/nats.go"
	"log"
)

type NatsSub struct {
	Conn      *nats.Conn
	Jm        NATSManager
	OffClient customer2.OfficeServiceClient
	UsClient  customer2.UserServiceClient
}

func NewNATS(jm NATSManager,
	OffClient customer2.OfficeServiceClient,
	UsClient customer2.UserServiceClient,
) *NatsSub {
	nc, err := nats.Connect(nats.DefaultURL) //"nats:4222")
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}

	return &NatsSub{
		Conn:      nc,
		Jm:        jm,
		OffClient: OffClient,
		UsClient:  UsClient,
	}
}

type NATSManager interface {
	DataSaveService(order *types.OrderRequest, offices []*types.Office, users []*types.User) error
}
