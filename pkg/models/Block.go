package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type Block struct {
	Number           int64          `json:"number" gorm:"primaryKey"`
	Hash             common.Hash    `json:"hash" gorm:"size:32;uniqueIndex"`
	ParentHash       common.Hash    `json:"parentHash" gorm:"size:32"`
	Nonce            string         `json:"nonce" gorm:"type:char(18)"`
	SHA3Uncles       common.Hash    `json:"sha3Uncles" gorm:"size:32"`
	LogsBloom        string         `json:"logsBloom" gorm:"size:514"`
	TransactionsRoot common.Hash    `json:"transactionsRoot" gorm:"size:32"`
	StateRoot        common.Hash    `json:"stateRoot" gorm:"size:32"`
	ReceiptsRoot     common.Hash    `json:"receiptsRoot" gorm:"size:32"`
	Miner            common.Address `json:"miner" gorm:"size:20"`
	ExtraData        []byte         `json:"extraData"`
	GasLimit         int64          `json:"gasLimit"`
	GasUsed          int64          `json:"gasUsed"`
	Timestamp        int64          `json:"timestamp"`
	Difficulty       string         `json:"difficulty"`
	// Uncles will be of type text or longtext depending on the driver type
	Uncles []common.Hash `json:"uncles" gorm:"serializer:stringer_joiner"`
	// Transactions will be of type text or longtext depending on the driver type
	Transactions  []common.Hash `json:"transactions" gorm:"serializer:stringer_joiner"`
	MixHash       common.Hash   `json:"mixHash" gorm:"size:32"`
	BaseFeePerGas *int64        `json:"baseFeePerGas"`
}

func (Block) TableName() string {
	return "blocks"
}
