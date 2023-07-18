package statrepository

import (
	"github.com/jmoiron/sqlx"
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
	return nil, nil
}
