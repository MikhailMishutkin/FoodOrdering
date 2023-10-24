package main

import (
	"log"

	"github.com/MikhailMishutkin/FoodOrdering/cmd"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	conf, err := configs.New("./configs/main.yaml")
	if err != nil {
		log.Fatalf("can't receive config data: %v\n", err)
	}
	err = app.StartGRPC(conf)
	if err != nil {
		log.Fatal(err)
	}
}
