package repository

import (
	"fmt"
	"os"

	"github.com/MikhailMishutkin/FoodOrdering/microservices/serializer"
	"github.com/jinzhu/copier"
	pb "gitlab.com/mediasoft-internship/final-task/contracts/pkg/contracts/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var fileBin string = "./product.bin"

func (r *RestaurantRepo) CreateProduct(p *pb.Product) error {

	file, err := os.OpenFile(fileBin, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	a := p.String()
	file.WriteString(a)
	fmt.Println(a)

	other := &pb.Product{}
	err = copier.Copy(other, p)
	if err != nil {
		return fmt.Errorf("cannot copy product data: %w", err)
	}
	//r.dataMap = make(map[string]*pb.Product)

	dataMap[other.Uuid] = other

	return nil
}

func (r *RestaurantRepo) GetProductList() (*pb.GetProductListResponse, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	_ = FromMapToSlice()

	p := new(pb.GetProductListResponse)

	err := serializer.ReadProtobufFromBinaryFile(fileBin, p)
	if err != nil {
		code := codes.Internal
		return nil, status.Errorf(code, "GetProductList went down witn error, cannot extract productlist from db: ", err)
	}

	return p, nil
}