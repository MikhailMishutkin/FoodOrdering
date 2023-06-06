package main

import (
	"log"

	app "github.com/MikhailMishutkin/FoodOrdering/cmd"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
)

func init() {

}

func main() {
	conf, err := configs.New("./configs/main.yaml.template")
	if err != nil {
		log.Fatalf("can't receive config data: %v\n", err)
	}
	gen.TypeSelector()
	gen.TypeSelector()

	if err := app.StartGRPCAndHTTPServer(conf); err != nil {
		log.Fatal(err)
	}

}
