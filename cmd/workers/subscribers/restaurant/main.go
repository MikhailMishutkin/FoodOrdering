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
	conf, err := configs.New("./configs/main.yaml")
	if err != nil {
		log.Fatalf("can't receive config data: %v\n", err)
	}

	if err = app.StartRestSubs(conf); err != nil {
		log.Fatal(err)
	}
}
