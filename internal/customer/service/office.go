package service

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
)

func (cu *CustomerUsecase) CreateOffice(office *types.Office) (err error) {
	log.Println("Service CreateOffice was invoked")
	err = cu.repoC.CreateOffice(office)
	return err
}

func (cu *CustomerUsecase) GetOfficeList() ([]*types.Office, error) {
	g, err := cu.repoC.GetOfficeList()
	return g, err
}
