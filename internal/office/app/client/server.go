package client

import (
	"context"
	"log"

	pb "github.com/MikhailMishutkin/FoodOrdering/pkg/contracts-v0.3.0/pkg/contracts/customer"

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

//TODO: переделать результат на GetActualMenuResponse, когда исправиться ситуация с v0.3.0
func GetActualMenu(g pb.OrderServiceClient) *pb.GetActualMenuResponse {
	log.Println("GetActualMenu was invoked")

	menu := &pb.GetActualMenuRequest{}
	res, err := g.GetActualMenu(context.Background(), menu)
	if err != nil {
		log.Fatalf("Error happend while getting actual menu: %v\n", err)
	}

	return res
}
