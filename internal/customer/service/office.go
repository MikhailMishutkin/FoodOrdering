package service

import (
	"fmt"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
)

func (r *CustomerUsecase) CreateOffice(office *pb.Office) (err error) {
	fmt.Println("Service CreateOffice was invoked")
	err = r.repoC.CreateOffice(office)
	return err
}

func (r *CustomerUsecase) GetOfficeList() ([]*pb.Office, error) {
	g, err := r.repoC.GetOfficeList()
	return g, err
}
