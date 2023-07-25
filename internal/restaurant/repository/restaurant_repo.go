package repository

import (
	"database/sql"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	//_ "github.com/jackc/pgx"
)

func RandomID() string {
	return uuid.New().String()
}

type RestaurantRepo struct {
	DB *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *RestaurantRepo {
	return &RestaurantRepo{
		DB: db,
	}
}

func NewDB() (*sql.DB, error) {

	c := configs.DB{}

	psqlInfo := fmt.Sprint(c.Conn)
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", "host=localhost port=5444 user=root password=root dbname=restaurant sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v\n", err)
	}

	return db, nil
}
