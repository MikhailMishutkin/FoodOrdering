package gormdb

import "time"

type Office struct {
	ID        int `gorm:"primary key;autoIncrement" json:"uuid"`
	Name      string
	Address   string
	CreatedAt time.Time
}
