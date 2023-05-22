package main

import (
	"log"

	"github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/app"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//config := configs.NewConfig()
	gen.TypeSelector()
	gen.TypeSelector()

	if err := app.StartGRPCRestaurantServer(); err != nil {
		log.Fatal(err)
	}
}
