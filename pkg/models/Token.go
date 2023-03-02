package models

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/datatypes"
)

type Token struct {
	Address     common.Address `gorm:"size:20;primaryKey"`
	Metadata    datatypes.JSON
	MetadataUrl string
	TokenID     string `gorm:"type:varchar(34);primaryKey"`
	Image       string
	Type        string
}

func (Token) TableName() string {
	return "eth_token"
}
