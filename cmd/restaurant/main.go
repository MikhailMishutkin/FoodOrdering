package main

import (
	"log"
	"os"

	"github.com/MikhailMishutkin/FoodOrdering/internal/app"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
)

func init() {
	os.Clearenv()
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
}

func main() {
	//config := configs.NewConfig()
	gen.TypeSelector()
	gen.TypeSelector()

	if err := app.StartGRPCAndHTTPServer(); err != nil {
		log.Fatal(err)
	}

}
