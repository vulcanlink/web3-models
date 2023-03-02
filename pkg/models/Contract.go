package models

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/datatypes"
)

type Contract struct {
	Address common.Address `gorm:"size:20;primaryKey"`
	// Foreign Key into "eth_code"
	CodeID       *common.Hash `gorm:"column:code;size:32"`
	Code         *Code
	Metadata     datatypes.JSON
	MetadataUri  string
	CompositeAbi datatypes.JSON
}

func (Contract) TableName() string {
	return "eth_contract"
}
