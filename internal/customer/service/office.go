package service

import (
	"fmt"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
)

func (cu *CustomerUsecase) CreateOffice(office *pb.Office) (err error) {
	fmt.Println("Service CreateOffice was invoked")
	err = cu.repoC.CreateOffice(office)
	return err
}

func (cu *CustomerUsecase) GetOfficeList() ([]*pb.Office, error) {
	g, err := cu.repoC.GetOfficeList()
	return g, err
}
