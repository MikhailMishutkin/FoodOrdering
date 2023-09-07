package repository

import (
	"github.com/jackc/pgx/v5"
	"log"

	"github.com/MikhailMishutkin/FoodOrdering/configs"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nats-io/nats.go"
)

type RestaurantRepo struct {
	DB   *pgx.Conn
	Conn *nats.Conn
}

func NewRestaurantRepo(db *pgx.Conn, conf configs.Config) *RestaurantRepo {
	nc, err := nats.Connect(conf.NATS.Host)
	if err != nil {
		log.Println("can't connect to NATS-server: %v", err)
	}

	return &RestaurantRepo{
		DB:   db,
		Conn: nc,
	}
}

//func NewDB() (*sql.DB, error) {
//
//	c, err := configs.New("./configs/main.yaml.template")
//	if err != nil {
//		return nil, fmt.Errorf("Can't load config in restaurant repo: %v\n", err)
//	}
//	psqlInfo := fmt.Sprint(c.DB.ConnSql)
//
//	db, err := sql.Open("postgres", psqlInfo)
//	if err != nil {
//		return nil, fmt.Errorf("can't connect to db: %v\n", err)
//	}
//
//	return db, nil
//}
