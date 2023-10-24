package bootstrap

import (
	"context"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/gormdb"
	_ "github.com/golang-migrate/migrate/v4/source"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*pgx.Conn, error) {
	c, err := configs.New("./configs/main.yaml")
	if err != nil {
		return nil, fmt.Errorf("Can't load config in restaurant repo: %v\n", err)
	}

	psqlInfo := fmt.Sprint(c.DB.ConnSql)

	db, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %v\n", err)
	}
	//defer db.Close(context.Background())
	//	defer db.Close()

	//m, err := migrate.New(
	//	"file://../FoodOrdering/migrations/restaurant/postgres",
	//	"postgres://root:root@restaurant-db:5432/restaurant?x-migrations-table=migrate.schema_migrations?sslmode=disable",
	//)
	//if err != nil {
	//	log.Println(err)
	//	return db, fmt.Errorf("can't automigrate: %v\n", err)
	//}
	//if err := m.Up(); err != nil {
	//	log.Println(err)
	//	fmt.Errorf("%v\n", err)
	//}

	return db, nil
}

func NewGormDB() (*gorm.DB, error) {

	c, err := configs.New("./configs/main.yaml")
	if err != nil {
		return nil, fmt.Errorf("Can't load config in restaurant repo: %v\n", err)
	}

	psqlInfo := fmt.Sprint(c.DB.ConnGorm)
	db, err := gorm.Open(postgres.Open(psqlInfo))
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
