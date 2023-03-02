package models

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

// NOTE: Pointers => nullable field, except for strings
type Transaction struct {
	Hash             common.Hash     `json:"hash" gorm:"size:32;primaryKey"`
	BlockHash        common.Hash     `json:"blockHash" gorm:"size:32"`
	BlockNumber      int64           `json:"blockNumber"`
	TransactionIndex int64           `json:"transactionIndex"`
	From             common.Address  `json:"from" gorm:"size:20"`
	To               *common.Address `json:"to" gorm:"size:20"`
	ContractAddress  *common.Address `json:"contractAddress" gorm:"size:20"`
	Value            string          `json:"value"`
	GasPrice         int64           `json:"gasPrice"`
	Gas              int64           `json:"gas"`
	GasUsed          int64           `json:"gasUsed"`
	Input            []byte          `json:"input"`
	Nonce            int64           `json:"nonce"`
	Logs             []*Log          `json:"logs"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	type TransactionAlias Transaction
	return json.Marshal(&struct {
		*TransactionAlias
		Logs LogsAsArrays `json:"logs"`
	}{
		TransactionAlias: (*TransactionAlias)(t),
		Logs:             LogsToLogsAsArrays(t.Logs),
	})
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type TransactionAlias Transaction
	aux := &struct {
		*TransactionAlias
		Logs LogsAsArrays `json:"logs"`
	}{
		TransactionAlias: (*TransactionAlias)(t),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	t.Logs = LogsAsArraysToLogs(aux.Logs)

	return nil
}
