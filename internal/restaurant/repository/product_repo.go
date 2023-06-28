package repository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"

	"log"
	"time"

	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/restaurant"
	_ "github.com/lib/pq"
)

func (r *RestaurantRepo) CreateProduct(b []byte) error {
	ndb, _ := NewDB()
	rr := NewRestaurantRepo(ndb)

	p := &pb.Product{}
	err := proto.Unmarshal(b, p)
	if err != nil {
		return fmt.Errorf("can't UNMARSHAL protodata to Product in repository CreateProduct: %v\n", err)
	}
	//fmt.Println(rr.db)
	tp := convertPbProdToProd(p)
	if err := rr.db.QueryRow(
		"INSERT INTO product (name, description, type_id, weight, price, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid",
		tp.Name,
		tp.Descript,
		tp.Type,
		tp.Weight,
		tp.Price,
		tp.CreatedAt,
	).Scan(&p.Uuid); err != nil {
		return err
	}

	return nil
}

func (r *RestaurantRepo) GetProductList() ([]byte, error) {
	p := &pb.Product{}
	tp := &types.Product{}
	ndb, _ := NewDB()
	rr := NewRestaurantRepo(ndb)
	products, err := rr.db.Query("SELECT * FROM product")
	if err != nil {
		log.Printf("Error to get groupnames from db: %s", err)
		return nil, err
	}
	defer products.Close()

	ps := []*pb.Product{}
	for products.Next() {
		if err = products.Scan(&tp.Uuid, &tp.Name, &tp.Descript, &tp.Type, &tp.Weight, &tp.Price, &tp.CreatedAt); err != nil {
			log.Printf("trouble with row.Next getproductlist: %v\n", err)
		}
		fmt.Println(tp)
		p = convertProdToPbProd(tp)
		ps = append(ps, p)
	}

	pp := &pb.GetProductListResponse{
		Result: ps,
	}
	protoData, err := proto.Marshal(pp)

	return protoData, nil
}

func convertProdToPbProd(tp *types.Product) *pb.Product {
	id := strconv.Itoa(tp.Uuid)
	t := timestamppb.New(tp.CreatedAt)

	pbp := &pb.Product{
		Uuid:        id,
		Name:        tp.Name,
		Description: tp.Descript,
		Type:        enumSelect(tp.Type),
		Weight:      int32(tp.Weight),
		Price:       tp.Price,
		CreatedAt:   t,
	}
	return pbp
}

func convertPbProdToProd(p *pb.Product) *types.Product {
	t := time.Unix(p.CreatedAt.Seconds, int64(p.CreatedAt.Nanos))
	tp := &types.Product{
		Name:      p.Name,
		Descript:  p.Description,
		Type:      int(p.Type.Number()),
		Weight:    int(p.Weight),
		Price:     p.Price,
		CreatedAt: t,
	}
	return tp
}

func enumSelect(i int) pb.ProductType {
	switch i {
	case 1:
		return pb.ProductType_PRODUCT_TYPE_SALAD
	case 2:
		return pb.ProductType_PRODUCT_TYPE_GARNISH
	case 3:
		return pb.ProductType_PRODUCT_TYPE_MEAT
	case 4:
		return pb.ProductType_PRODUCT_TYPE_SOUP
	case 5:
		return pb.ProductType_PRODUCT_TYPE_DRINK
	case 6:
		return pb.ProductType_PRODUCT_TYPE_DESSERT
	default:
		return pb.ProductType_PRODUCT_TYPE_UNSPECIFIED
	}

}

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
