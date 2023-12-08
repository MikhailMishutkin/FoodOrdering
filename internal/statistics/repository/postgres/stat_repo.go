package statrepository

import (
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"log"
)

type StatRepo struct {
	DB   *sqlx.DB
	Conn *nats.Conn
}

func NewStatRepo(db *sqlx.DB, conf configs.Config) *StatRepo {
	nc, err := nats.Connect(conf.NATS.Host)
	if err != nil {
		log.Println("can't connect to NATS-server: %v", err)
	}

	//TODO
	log.Println(db)

	return &StatRepo{
		DB:   db,
		Conn: nc,
	}
}
