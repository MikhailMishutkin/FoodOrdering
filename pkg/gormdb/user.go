package gormdb

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID        int //`gorm:"primary key;autoIncrement" json:"uuid"`
	Name      string
	OfficeID  int
	CreatedAt time.Time
	Office    Office `gorm:"foreignKey:OfficeID"`
}
