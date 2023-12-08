package types

import "time"

type MenuCreate struct {
	OnDate    time.Time
	OpenAt    time.Time
	ClosedAt  time.Time
	Salads    []string
	Garnishes []string
	Meats     []string
	Soups     []string
	Drinks    []string
	Desserts  []string
}

type Menu struct {
	Uuid      int
	OnDate    time.Time
	OpenAt    time.Time
	ClosedAt  time.Time
	Salads    []*Product
	Garnishes []*Product
	Meats     []*Product
	Soups     []*Product
	Drinks    []*Product
	Desserts  []*Product
	CreatedAt time.Time
}
