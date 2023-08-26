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

//func NewGormDB() (*gorm.DB, error) {
//	c, err := configs.New("./configs/main.yaml.template")
//	if err != nil {
//		return nil, err
//	}
//	psqlInfo := fmt.Sprint(c.DB.ConnGorm)
//
//	db, err := gorm.Open(postgres.Open(psqlInfo))
//	if err != nil {
//		return nil, fmt.Errorf("can't connect to db: %v\n", err)
//	}
//
//	err = db.AutoMigrate(&gormdb.Office{})
//	if err != nil {
//		return nil, fmt.Errorf("cannot create table office in gorm: %v\n", err)
//	}
//
//	err = db.AutoMigrate(&gormdb.User{})
//	if err != nil {
//		return nil, fmt.Errorf("cannot create table users in gorm: %v\n", err)
//	}
//
//	return db, nil
//}
