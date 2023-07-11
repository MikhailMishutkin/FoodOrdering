package types

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
	Count       int
	ProductUuid int
}
