package repository

import (
	"io/ioutil"
	"log"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/protobuf/proto"
)

func DateConv(t time.Time) time.Time {
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
