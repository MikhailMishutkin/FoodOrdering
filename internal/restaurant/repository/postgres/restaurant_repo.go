package repository

import (
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type RestaurantRepo struct {
	DB *pgx.Conn
}

func NewRestaurantRepo(db *pgx.Conn) *RestaurantRepo {
	return &RestaurantRepo{
		DB: db,
	}
}
