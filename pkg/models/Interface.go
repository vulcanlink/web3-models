package models

import (
	"gorm.io/datatypes"
)

type Interface struct {
	ID   string `gorm:"primaryKey;type:char(10)"`
	Name string
	Abi  datatypes.JSON
}

func (Interface) TableName() string {
	return "eth_interface"
}
