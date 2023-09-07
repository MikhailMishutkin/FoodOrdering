package statrepository

import (
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type StatRepo struct {
	DB            *sqlx.DB
	ClientProduct restaurant.ProductServiceClient
	ClientMenu    restaurant.MenuServiceClient
	Conn          *nats.Conn
}

func NewStatRepo(db *sqlx.DB, conf configs.Config) *StatRepo {
	nc, err := nats.Connect(conf.NATS.Host)
	if err != nil {
		log.Println("can't connect to NATS-server: %v", err)
	}
	conn, err := grpc.Dial(conf.API.GHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	//TODO
	log.Println(db)

	return &StatRepo{
		DB:            db,
		ClientProduct: restaurant.NewProductServiceClient(conn),
		ClientMenu:    restaurant.NewMenuServiceClient(conn),
		Conn:          nc,
	}
}

//func NewDB() (*sqlx.DB, error) {
//	db, err := sqlx.Connect(
//		"postgres",
//		"host=localhost port=5446 user=root password=root dbname=statistics sslmode=disable",
//	)
//
//	if err != nil {
//		return nil, fmt.Errorf("can't connect to db statistcs: %v\n", err)
//	}
//	return db, err
//}
