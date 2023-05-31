package client

import (
	"context"
	"log"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/contracts-v0.3.0/pkg/contracts/restaurant"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var addr string = "localhost:5051"

func Сonn() (*pb.GetMenuResponse, error) {
	log.Println("Conn was invoked")

	// tls := true // change that to false if needed
	// opts := []grpc.DialOption{}

	// if tls {
	// 	certFile := "ssl/ca.crt"
	// 	creds, err := credentials.NewClientTLSFromFile(certFile, "")

	// 	if err != nil {
	// 		log.Fatalf("Error while loading CA trust certificate: %v\n", err)
	// 	}

	// 	opts = append(opts, grpc.WithTransportCredentials(creds))
	// }

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	t := time.Now()
	t1 := t.AddDate(0, 0, 1)
	ts := timestamppb.New(t1)
	//fmt.Println(ts)

	c := pb.NewMenuServiceClient(conn)

	mr := &pb.GetMenuRequest{
		OnDate: ts,
	}
	// fmt.Println(mr)

	res, err := c.GetMenu(context.Background(), mr)
	if err != nil {
		log.Println("Can't get the menu from restaurant", err)
		return nil, err
	}

	return res, nil

}

// func GetActualMenu(g pb.OrderServiceClient) *pb.GetActualMenuResponse {
// 	log.Println("GetActualMenu was invoked")

// 	req := &pb.GetActualMenuRequest{}
// 	res, err := g.GetActualMenu(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error happend while getting actual menu: %v\n", err)
// 	}

// 	return res
// }
