package types

import "time"

type Product struct {
	Uuid      int       `db:"id"`
	Name      string    `db:"name"`
	Descript  string    `db:"descr"`
	Type      int       `db:"type"`
	Weight    int       `db:"weight"`
	Price     float64   `db:"price"`
	CreatedAt time.Time `db:"created_at"`
}

type StatProduct struct {
	Uuid  int
	Name  string
	Count int
	Type  int
}
