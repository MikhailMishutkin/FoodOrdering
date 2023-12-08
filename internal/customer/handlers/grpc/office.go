package handlerscustomer

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/microservices/gen"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/customer"
	"log"
)

func (s *CustomerService) CreateOffice(ctx context.Context, in *customer.CreateOfficeRequest) (*customer.CreateOfficeResponse, error) {
	log.Print("CreteOffice was invoked")
	office := &types.Office{}
	var err error
	if in.Name == "" && in.Address == "" {
		office = gen.NewOffice()
	} else {
		office.Name = in.Name
		office.Address = in.Address
		err = s.cs.CreateOffice(office)
		if err != nil {
			return nil, err
		}
	}

	return &customer.CreateOfficeResponse{}, err
}

func (s *CustomerService) GetOfficeList(context.Context, *customer.GetOfficeListRequest) (*customer.GetOfficeListResponse, error) {
	res, err := s.cs.GetOfficeList()
	if err != nil {
		return nil, err
	}

	resCust := convertOffice(res)
	r := &customer.GetOfficeListResponse{
		Result: resCust,
	}

	return r, err
}
