package cusrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
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

//TODO: connect db
func (cr *CustomerRepo) CreateOrder(order *pb.CreateOrderRequest) error {
	fmt.Println("save neworder in db: ", order)
	conf, err := configs.New("./configs/main.yaml.template")
	if err != nil {
		log.Println("can't load config (repository rest/order: ", err)
	}

	data, err := proto.Marshal(order)
	if err != nil {
		fmt.Errorf("cannot marshal proto message to binary: %w", err)
		return err
	}
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("can't connect to NATS: %v", err)
		return err
	}
	defer nc.Drain()

	js, err := nc.JetStream()
	if err != nil {
		log.Println("error line 141 app: ", err)
	}

	js.AddStream(&nats.StreamConfig{
		Name:     conf.NATS.Name,
		Subjects: []string{"orders"},
	})

	//conf.Storage = nats.FileStorage?????

	ack, err := js.Publish("orders", data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ack)
	return nil
}

//func processMsg(sub *nats.Subscription) {
//	for sub.IsValid() {
//		msgs, err := sub.Fetch(1)
//		if err == nil {
//			fmt.Printf("INFO - Got reply - %s\n", string(msgs[0].Data))
//		}
//	}
//}
