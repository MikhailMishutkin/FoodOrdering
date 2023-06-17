package cusrepository

import (
	"fmt"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/google/uuid"
	"sync"
)

var officeMap map[string]*pb.Office

func init() {
	officeMap = make(map[string]*pb.Office)
}
func RandomID() string {
	return uuid.New().String()
}

type CustomerRepo struct {
	mutex     sync.RWMutex
	officeMap map[string]*pb.Office
}

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{
		officeMap: officeMap,
	}
}

func (cr *CustomerRepo) CreateOrder(order *pb.CreateOrderRequest) error {
	fmt.Println("save neworder in db", order)
	//TODO: connect db
	return nil
}
