package main

import (
	"log"

	restaurant_app "github.com/MikhailMishutkin/FoodOrdering/cmd/restaurant/restaurant_app"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
)

func init() {
	//os.Clearenv()
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
}

func main() {
	//config := configs.NewConfig()
	gen.TypeSelector()
	gen.TypeSelector()

	if err := restaurant_app.StartGRPCAndHTTPServer(); err != nil {
		log.Fatal(err)
	}

}
