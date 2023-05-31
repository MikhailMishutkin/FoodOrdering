package main

import "github.com/MikhailMishutkin/FoodOrdering/internal/app"

// import (
// 	"log"

// 	cs "github.com/MikhailMishutkin/FoodOrdering/internal/customer/app/client"
// )

// // func init() {
// // 	err := godotenv.Load()
// // 	if err != nil {
// // 		log.Fatal("Error loading .env file")
// // 	}
// // }

// func main() {
// 	//config := configs.NewConfig()
// 	// gen.TypeSelector()
// 	// gen.TypeSelector()

// 	if err := cs.StartGRPCCustomerServer(); err != nil {
// 		log.Fatal(err)
// 	}
// }
func main() {
	app.StartGRPC()
}
