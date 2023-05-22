package repository

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	pb "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/protobuf/proto"
)

func dateConv(t time.Time) time.Time {
	var layout = "02.01.2006"
	t1 := t.Format(layout)

	tConv, _ := time.Parse(layout, t1)

	return tConv
}

func FromMapToSlice() (sp []*pb.Product) {

	for _, v := range dataMap {
		p := &pb.Product{}
		p = v
		sp = append(sp, p)
	}
	gplr := new(pb.GetProductListResponse)
	gplr.Result = sp
	out, err := proto.Marshal(gplr)
	if err != nil {
		log.Fatalln("Failed to encode product:", err)
	}
	if err := ioutil.WriteFile(fileBin, out, 0644); err != nil {
		log.Fatalln("Failed to write product:", err)
	}

	return sp
}

func NatsPublisher() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	for {
		nc.Publish("intros", []byte("Hello1 World!"))
		time.Sleep(1 * time.Second)
	}
}
