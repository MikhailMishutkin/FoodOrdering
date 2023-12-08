package cusrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/gormdb"
	"log"
)

func (cr *CustomerRepo) CreateOffice(office *types.Office) error {
	fmt.Println("Repo CreateOffice was invoked")

	gOffice := &gormdb.Office{
		Name:    office.Name,
		Address: office.Address,
	}

	err := cr.DB.Create(gOffice).Error
	if err != nil {
		return fmt.Errorf("can't create office in gorm: %v\n", err)
	}
	return err
}

func (cr *CustomerRepo) GetOfficeList() ([]*types.Office, error) {
	log.Printf("GetOfficeList Repository was invoked")
	officesList := &[]gormdb.Office{}
	err := cr.DB.Find(officesList).Error
	if err != nil {
		return nil, fmt.Errorf("can't create office in gorm: %v\n", err)
	}
	tOffs := []*types.Office{}

	for _, v := range *officesList {
		tOff := &types.Office{}
		tOff.Uuid = v.ID
		tOff.Name = v.Name
		tOff.Address = v.Address
		tOff.CreatedAt = v.CreatedAt
		tOffs = append(tOffs, tOff)
	}
	return tOffs, nil
}
