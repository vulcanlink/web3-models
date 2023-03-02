package models

import "github.com/ethereum/go-ethereum/common"

type SupportsInterface struct {
	Address common.Address `gorm:"size:20;primaryKey"`
	// Foreign Key into "eth_code"
	Interface string `gorm:"primaryKey;type:char(10)"`
	Supports  bool
}

func (SupportsInterface) TableName() string {
	return "eth_supports_interface"
}
