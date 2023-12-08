package cusrepository

import (
	gorm "gorm.io/gorm"
)

type CustomerRepo struct {
	DB *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{
		DB: db,
	}
}
