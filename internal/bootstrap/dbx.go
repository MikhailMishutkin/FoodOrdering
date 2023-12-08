package bootstrap

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/configs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDBX() (*sqlx.DB, error) {
	c, err := configs.New("./configs/main.yaml")
	if err != nil {
		return nil, fmt.Errorf("Can't load config in restaurant repo: %v\n", err)
	}
	psqlInfo := fmt.Sprint(c.DB.ConnSqlx)
	db, err := sqlx.Connect(
		"postgres",
		psqlInfo,
	)

	if err != nil {
		return nil, fmt.Errorf("can't connect to db statistcs: %v\n", err)
	}
	return db, err
}
