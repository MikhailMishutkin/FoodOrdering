package cusrepository

import (
	"fmt"
	"log"
	"sync"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/contracts-v0.3.0/pkg/contracts/customer"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

var dataMap map[string]*pb.Product

func init() {
	dataMap = make(map[string]*pb.Product)
}
func RandomID() string {
	return uuid.New().String()
}

type CustomerRepo struct {
	mutex   sync.RWMutex
	dataMap map[string]*pb.Product
}

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{
		dataMap: dataMap,
	}
}

func natsSubscriber() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	nc.Subscribe("intros", func(m *nats.Msg) {
		fmt.Println(string(m.Data))
	})
	time.Sleep(30 * time.Second)

}

func (cr *CustomerRepo) GetActualMenu(menu *pb.GetActualMenuResponse) {
	m := dataMap
	salads := menu.Salads
	for _, v := range salads {
		m[v.Uuid] = v
	}

	//go natsSubscriber()
}

func (cr *CustomerRepo) CreateOrder() {

}
