package statrepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	//_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type StatRepo struct {
	DB *sqlx.DB
}

func NewStatRepo(db *sqlx.DB) *StatRepo {
	return &StatRepo{
		DB: db,
	}
}

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		"host=localhost port=5446 user=root password=root dbname=statistics sslmode=disable",
	)
	if err != nil {
		return nil, fmt.Errorf("can't connect to db statistcs: %v\n", err)
	}
	return db, err
}
