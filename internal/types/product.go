package types

import "time"

type Product struct {
	Uuid      int       `db:"uuid"`
	Name      string    `db:"name"`
	Descript  string    `db:"description"`
	Type      int       `db:"type_id"`
	Weight    int       `db:"weight"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}

type StatProduct struct {
	Uuid  int    `db:"product_id"`
	Name  string `db:"name"`
	Count int    `db:"sum"`
	Type  int    `db:"type_id"`
}
