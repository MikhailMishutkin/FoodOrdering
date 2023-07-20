package types

import "time"

type OrderRequest struct {
	UserUuid  int
	Salads    []*OrderItem
	Garnishes []*OrderItem
	Meats     []*OrderItem
	Soups     []*OrderItem
	Drinks    []*OrderItem
	Desserts  []*OrderItem
}

type OrderItem struct {
	ID          int       `db:"id"'`
	OnDate      time.Time `db:"on_date"`
	ProductUuid int       `db:"product_id"'`
	Count       int       `db:"count"`
	ProductName string
}

type OrderByOffice struct {
	OfficeUuid    int
	OfficeName    string
	OfficeAddress string
	Result        []*OrderItem
}
