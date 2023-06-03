package customer_app

import (
	"log"
	"net"

	handlers "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/handlers/grpc"
	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository"
	rest "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	"google.golang.org/grpc"
)

var gAddr string = "0.0.0.0:5051"

func StartGRPC() {
	repo := repository.NewRestaurantRepo()
	// log.Print("check repo: ", repo)
	rs := handlers.NewRestaurantService(repo)
	lis, err := net.Listen("tcp", gAddr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", gAddr)

	s := grpc.NewServer()

	rest.RegisterMenuServiceServer(s, rs)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
