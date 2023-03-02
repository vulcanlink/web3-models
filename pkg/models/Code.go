package models

import (
	"database/sql"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/datatypes"
)

type Code struct {
	// CodeHash common.Hash `gorm:"primaryKey"`
	ID             common.Hash `gorm:"column:code_hash;size:32;primaryKey"`
	Code           string
	Abi            datatypes.JSON
	AbiLastChecked sql.NullTime
}

func (Code) TableName() string {
	return "eth_code"
}
