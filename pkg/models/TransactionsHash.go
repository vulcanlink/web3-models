package models

import "github.com/ethereum/go-ethereum/common"

type TransactionsHash struct {
	Hash common.Hash `json:"hash" gorm:"size:32"`
}

func (TransactionsHash) TableName() string {
	return "transactions_hash"
}
