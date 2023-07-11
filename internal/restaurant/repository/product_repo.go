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

//
//func convertProdToPbProd(tp *types.Product) *pb.Product {
//	id := strconv.Itoa(tp.Uuid)
//	t := timestamppb.New(tp.CreatedAt)
//
//	pbp := &pb.Product{
//		Uuid:        id,
//		Name:        tp.Name,
//		Description: tp.Descript,
//		Type:        enumSelect(tp.Type),
//		Weight:      int32(tp.Weight),
//		Price:       tp.Price,
//		CreatedAt:   t,
//	}
//	return pbp
//}
//
//func convertPbProdToProd(p *pb.Product) *types.Product {
//	//t := time.Unix(p.CreatedAt.Seconds, int64(p.CreatedAt.Nanos))
//	tp := &types.Product{
//		Name:      p.Name,
//		Descript:  p.Description,
//		Type:      int(p.Type.Number()),
//		Weight:    int(p.Weight),
//		Price:     p.Price,
//		CreatedAt: time.Now(),
//	}
//	return tp
//}

// при работе без базы данных
//Create
//file, err := os.OpenFile(fileBin, os.O_APPEND|os.O_WRONLY, 0666)
//if err != nil {
//	fmt.Println(err)
//}
//defer file.Close()
//
//a := p.String()
//file.WriteString(a)
//fmt.Println(a)
//
//other := &pb.Product{}
//err = copier.Copy(other, p)
//if err != nil {
//	return fmt.Errorf("cannot copy product data: %w", err)
//}
//dataMap := make(map[string]*pb.Product)
//dataMap[other.Uuid] = other

//ProductList
//r.mutex.RLock()
//defer r.mutex.RUnlock()
//
//_ = FromMapToSlice()
//
//p := new(pb.GetProductListResponse)
//
//err := serializer.ReadProtobufFromBinaryFile(fileBin, p)
//if err != nil {
//	code := codes.Internal
//	return nil, status.Errorf(code, "GetProductList went down witn error, cannot extract productlist from db: %v\n", err)
//}
