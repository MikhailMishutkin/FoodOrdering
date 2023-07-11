package main

import (
	app "github.com/MikhailMishutkin/FoodOrdering/cmd"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {
	conf, err := configs.New("./configs/main.yaml.template")
	if err != nil {
		log.Fatalf("can't receive config data: %v\n", err)
	}

	if err := app.StartGRPCAndHTTPServer(conf); err != nil {
		log.Fatal(err)
	}

}
