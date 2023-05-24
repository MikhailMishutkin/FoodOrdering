package client

import (
	"context"
	"log"

	pb "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func conn() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)

	GetActualMenu(c)
}

func GetActualMenu(g pb.OrderServiceClient) {
	log.Println("GetActualMenu was invoked")

	menu := &pb.GetActualMenuRequest{}
	res, err := g.GetActualMenu(context.Background(), menu)
}
