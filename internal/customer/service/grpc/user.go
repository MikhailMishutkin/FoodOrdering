package service

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
)

func (cu *CustomerUsecase) CreateUser(user *types.User) error {
	err := cu.repoC.CreateUser(user)
	return err
}

func (cu *CustomerUsecase) GetUserList(ofId int) ([]*types.User, error) {
	res, err := cu.repoC.GetUserList(ofId)
	return res, err
}
