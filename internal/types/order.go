package types

type OrderRequest struct {
	UserUuid  string
	Salads    []*OrderItem
	Garnishes []*OrderItem
	Meats     []*OrderItem
	Soups     []*OrderItem
	Drinks    []*OrderItem
	Desserts  []*OrderItem
}

type OrderItem struct {
	Count       int32
	ProductUuid string
}
