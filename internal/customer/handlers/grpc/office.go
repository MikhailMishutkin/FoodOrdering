package handlers_customer

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
	"github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
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

func convertOffice(res []*types.Office) []*customer.Office {
	var resPb []*customer.Office

	for _, v := range res {
		id := strconv.Itoa(v.Uuid)
		t := timestamppb.New(v.CreatedAt)
		pr := &customer.Office{
			Uuid:      id,
			Name:      v.Name,
			Address:   v.Address,
			CreatedAt: t,
		}
		resPb = append(resPb, pr)
	}
	return resPb
}
