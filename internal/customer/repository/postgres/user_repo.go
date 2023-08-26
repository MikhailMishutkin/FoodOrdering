package cusrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/gormdb"
	"log"
)

func (cr *CustomerRepo) CreateUser(user *types.User) error {
	log.Println("CreateUser repo was invoked")
	gUser := &gormdb.User{
		Name:     user.Name,
		OfficeID: user.OfficeUuid,
	}

	err := cr.DB.Create(gUser).Error
	if err != nil {
		return fmt.Errorf("can't create office in gorm: %v\n", err)
	}
	return err
	return nil
}

func (cr *CustomerRepo) GetUserList(id int) ([]*types.User, error) {
	log.Printf("GetOfficeList Repository was invoked")
	usersList := &[]gormdb.User{}
	err := cr.DB.Where("office_id = ?", id).Find(usersList).Error
	if err != nil {
		return nil, fmt.Errorf("can't create office in gorm: %v\n", err)
	}

	tUsers := []*types.User{}

	for _, v := range *usersList {
		gOff := &gormdb.Office{}
		err = cr.DB.Where("id = ?", id).First(gOff).Error
		if err != nil {
			return nil, fmt.Errorf("can't find officeid in gorm (GetUserList): %v\n", err)
		}
		tUser := &types.User{}
		tUser.Uuid = v.ID
		tUser.Name = v.Name
		tUser.OfficeUuid = v.OfficeID
		tUser.OfficeName = gOff.Name
		tUser.CreatedAt = v.CreatedAt
		tUsers = append(tUsers, tUser)
	}

	return tUsers, err
}
