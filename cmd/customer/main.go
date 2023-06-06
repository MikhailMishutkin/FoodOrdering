package main

import (
	"log"

	"github.com/MikhailMishutkin/FoodOrdering/cmd"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
)

//import "github.com/MikhailMishutkin/FoodOrdering/cmd/customer/customer_app"

func main() {
	conf, err := configs.New("./configs/main.yaml.template")
	if err != nil {
		log.Fatalf("can't receive config data: %v\n", err)
	}
	app.StartGRPC(conf)
}
