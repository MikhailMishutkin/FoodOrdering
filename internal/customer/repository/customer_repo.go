package cusrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/gormdb"
	"gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"log"
)

type CustomerRepo struct {
	DB *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{
		DB: db,
	}
}

func NewGormDB() (*gorm.DB, error) {
	c := configs.DB{}

	psqlInfo := fmt.Sprint(c.Conn)
	fmt.Println(psqlInfo)

	db, err := gorm.Open(postgres.Open("host=localhost port=5445 user=root password=root dbname=customer sslmode=disable"))
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v\n", err)
	}

	err = db.AutoMigrate(&gormdb.Office{})
	if err != nil {
		log.Fatalf("cannot create tables in gorm: %v\n", err)
	}

	err = db.AutoMigrate(&gormdb.User{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
