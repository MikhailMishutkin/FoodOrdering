package types

import "time"

type Product struct {
	Uuid      int
	Name      string
	Descript  string
	Type      int
	Weight    int
	Price     float64
	CreatedAt time.Time
}

type StatProduct struct {
	Uuid  int
	Name  string
	Count int
	Type  int
}
