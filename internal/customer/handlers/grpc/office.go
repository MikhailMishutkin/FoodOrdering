package handlers_customer

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
)

func (s *CustomerService) CreateOffice(ctx context.Context, in *customer.CreateOfficeRequest) (*customer.CreateOfficeResponse, error) {
	var office *customer.Office
	if in.Name == "" && in.Address == "" {
		office = gen.NewOffice()
	} else {
		office.Name = in.Name
		office.Address = in.Address
		office.Uuid = gen.RandomID()
	}

	err := s.cs.CreateOffice(office)

	return &customer.CreateOfficeResponse{}, err
}

func (s *CustomerService) GetOfficeList(context.Context, *customer.GetOfficeListRequest) (*customer.GetOfficeListResponse, error) {
	res, err := s.cs.GetOfficeList()
	if err != nil {
		return nil, err
	}
	r := &customer.GetOfficeListResponse{
		Result: res,
	}
	return r, err
}
