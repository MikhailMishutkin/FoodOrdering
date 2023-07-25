package repository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"

	_ "github.com/lib/pq"
)

func (r *RestaurantRepo) CreateProduct(product *types.Product) error {

	if err := r.DB.QueryRow(
		"INSERT INTO product (name, description, type_id, weight, price) VALUES ($1, $2, $3, $4, $5) RETURNING uuid",
		product.Name,
		product.Descript,
		product.Type,
		product.Weight,
		product.Price,
	).Scan(&product.Uuid); err != nil {
		return err
	}

	return nil
}

func (r *RestaurantRepo) GetProductList() ([]*types.Product, error) {

	ps := make([]*types.Product, 0, 24)

	products, err := r.DB.Query("SELECT * FROM product")
	if err != nil {
		return nil, fmt.Errorf("Error to get ProductList from db: %s", err)

	}
	defer products.Close()

	for products.Next() {
		tp := &types.Product{}
		if err = products.Scan(&tp.Uuid, &tp.Name, &tp.Descript, &tp.Type, &tp.Weight, &tp.Price, &tp.CreatedAt); err != nil {
			return nil, fmt.Errorf("trouble with row.Next getproductlist: %v\n", err)
		}

		ps = append(ps, tp)

	}

	return ps, nil
}
