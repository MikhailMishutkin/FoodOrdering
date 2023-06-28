package repository

import (
	"database/sql"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"sync"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

//var dataMap map[string]*pb.Product

func RandomID() string {
	return uuid.New().String()
}

type RestaurantRepo struct {
	mutex sync.RWMutex
	//dataMap map[string]*pb.Product
	db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{
		db: db,
	}
}

func NewDB() (*sql.DB, error) {

	c := configs.DB{}

	psqlInfo := fmt.Sprint(c.Conn)
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", "user=root password=123 dbname=restaurant sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v\n", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("no ping by db: %v\n", err)
	}
	return db, nil
}
