package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/gormdb"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*sql.DB, error) {
	c, err := configs.New("./configs/main.yaml.template")
	if err != nil {
		return nil, fmt.Errorf("Can't load config in restaurant repo: %v\n", err)
	}

	psqlInfo := fmt.Sprint(c.DB.ConnSql)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v\n", err)
	}

	m, err := migrate.New("file://migrations/restaurant/postgres", "postgres://root:root@localhost:5444/restaurant?sslmode=disable")
	if err != nil {
		fmt.Errorf("can't automigrate: %v\n", err)
	}
	if err := m.Up(); err != nil {
		fmt.Errorf("%v\n", err)
	}

	return db, nil
}

func NewGormDB() (*gorm.DB, error) {

	//c, err := configs.New("./configs/main.yaml.template")
	//if err != nil {
	//	return nil, err
	//}

	//psqlInfo := fmt.Sprint(c.DB.ConnGorm)

	db, err := gorm.Open(postgres.Open("port=5445 user=root password=root dbname=customer sslmode=disable"))
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v\n", err)
	}

	err = db.AutoMigrate(&gormdb.Office{})
	if err != nil {
		return nil, fmt.Errorf("cannot create table office in gorm: %v\n", err)
	}

	err = db.AutoMigrate(&gormdb.User{})
	if err != nil {
		return nil, fmt.Errorf("cannot create table users in gorm: %v\n", err)
	}

	return db, nil
}

func NewDBX() (*sqlx.DB, error) {
	db, err := sqlx.Connect(
		"postgres",
		"host=localhost port=5446 user=root password=root dbname=statistics sslmode=disable",
	)

	if err != nil {
		return nil, fmt.Errorf("can't connect to db statistcs: %v\n", err)
	}
	return db, err
}
