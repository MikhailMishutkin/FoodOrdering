package types

type Profit struct {
	Count int     `db:"count"`
	Price float64 `db:"price"`
}
