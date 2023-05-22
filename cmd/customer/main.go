package main

import (
	"log"

	so "github.com/MikhailMishutkin/FoodOrdering/internal/office/app/server"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func main() {
	//config := configs.NewConfig()
	// gen.TypeSelector()
	// gen.TypeSelector()

	if err := so.StartGRPCCustomerServer(); err != nil {
		log.Fatal(err)
	}
}
