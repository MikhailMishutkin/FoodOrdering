package cusrepository

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	pb "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
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

func (cr *CustomerRepo) GetMenu() {

	go natsSubscriber()
}