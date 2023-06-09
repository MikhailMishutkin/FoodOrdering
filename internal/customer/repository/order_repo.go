package cusrepository

import (
	"fmt"
	"log"
	"sync"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

var officeMap map[string]*pb.Office

func init() {
	officeMap = make(map[string]*pb.Office)
}
func RandomID() string {
	return uuid.New().String()
}

type CustomerRepo struct {
	mutex     sync.RWMutex
	officeMap map[string]*pb.Office
}

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{
		officeMap: officeMap,
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

func (cr *CustomerRepo) CreateOrder() {

}
