package cusrepository

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/nats-io/nats.go"
	"log"
)

// TODO: connect db
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

//func processMsg(sub *nats.Subscription) {

//	for sub.IsValid() {
//		msgs, err := sub.Fetch(1)
//		if err == nil {
//			fmt.Printf("INFO - Got reply - %s\n", string(msgs[0].Data))
//		}
//	}
//}

//without db
//var officeMap map[string]*pb.Office

//func init() {
//	officeMap = make(map[string]*pb.Office)
//}
//func RandomID() string {
//	return uuid.New().String()
//}

//with Jetstream
//conf, err := configs.New("./configs/main.yaml.template")
//if err != nil {
//	log.Println("can't load config (repository rest/order: ", err)
//}

//data, err := proto.Marshal(order)
//if err != nil {
//	fmt.Errorf("cannot marshal proto message to binary: %w", err)
//	return err
//}
//js, err := nc.JetStream()
//	if err != nil {
//		log.Println("error line 141 app: ", err)
//	}
//
//	js.AddStream(&nats.StreamConfig{
//		Name:     conf.NATS.Name,
//		Subjects: []string{"orders"},
//	})
//
//	//conf.Storage = nats.FileStorage?????
//
//	ack, err := js.Publish("orders", data)
//	if err != nil {
//		log.Println(err)
//	}
//	fmt.Println(ack)
