package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vulcanlink/web3-models/pkg/slicefp"
)

// Primary key: BlockNumber + LogIndex
type Log struct {
	BlockNumber      int64          `json:"blockNumber" gorm:"primaryKey"`
	LogIndex         int64          `json:"logIndex" gorm:"primaryKey"`
	BlockHash        common.Hash    `json:"blockHash" gorm:"size:32"`
	TransactionHash  common.Hash    `json:"transactionHash" gorm:"size:32"`
	TransactionIndex int64          `json:"transactionIndex"`
	Address          common.Address `json:"address" gorm:"size:20"`
	Data             []byte         `json:"data"`
	Topic0           common.Hash    `json:"topic0" gorm:"size:32;index:idx_721,priority:1;index:idx_1155,priority:1"`
	Topic1           common.Hash    `json:"topic1" gorm:"size:32"`
	Topic2           common.Hash    `json:"topic2" gorm:"size:32;index:idx_721"`
	Topic3           common.Hash    `json:"topic3" gorm:"size:32;index:idx_1155"`
	Removed          bool           `json:"removed"`
}

func (Log) TableName() string {
	return "eth_log"
}

type LogsAsArrays struct {
	BlockNumber      []int64          `json:"blockNumber" gorm:"primaryKey"`
	LogIndex         []int64          `json:"logIndex" gorm:"primaryKey"`
	BlockHash        []common.Hash    `json:"blockHash" gorm:"size:32"`
	TransactionHash  []common.Hash    `json:"transactionHash" gorm:"size:32"`
	TransactionIndex []int64          `json:"transactionIndex"`
	Address          []common.Address `json:"address" gorm:"size:20"`
	Data             [][]byte         `json:"data"`
	Topic0           []common.Hash    `json:"topic0" gorm:"size:32;index:idx_721,priority:1;index:idx_1155,priority:1"`
	Topic1           []common.Hash    `json:"topic1" gorm:"size:32"`
	Topic2           []common.Hash    `json:"topic2" gorm:"size:32;index:idx_721"`
	Topic3           []common.Hash    `json:"topic3" gorm:"size:32;index:idx_1155"`
	Removed          []bool           `json:"removed"`
}

func LogsToLogsAsArrays(logs []*Log) LogsAsArrays {
	numLogs := len(logs)

	if 0 == numLogs {
		return LogsAsArrays{}
	}

	l := LogsAsArrays{
		BlockNumber:      make([]int64, numLogs),
		LogIndex:         make([]int64, numLogs),
		BlockHash:        make([]common.Hash, numLogs),
		TransactionHash:  make([]common.Hash, numLogs),
		TransactionIndex: make([]int64, numLogs),
		Address:          make([]common.Address, numLogs),
		Data:             make([][]byte, numLogs),
		Topic0:           make([]common.Hash, numLogs),
		Topic1:           make([]common.Hash, numLogs),
		Topic2:           make([]common.Hash, numLogs),
		Topic3:           make([]common.Hash, numLogs),
		Removed:          make([]bool, numLogs),
	}

	slicefp.SliceEachCoroIdx(logs, func(log *Log, i int) {
		l.BlockNumber[i] = log.BlockNumber
		l.LogIndex[i] = log.LogIndex
		l.BlockHash[i] = log.BlockHash
		l.TransactionHash[i] = log.TransactionHash
		l.TransactionIndex[i] = log.TransactionIndex
		l.Address[i] = log.Address
		l.Data[i] = log.Data
		l.Topic0[i] = log.Topic0
		l.Topic1[i] = log.Topic1
		l.Topic2[i] = log.Topic2
		l.Topic3[i] = log.Topic3
		l.Removed[i] = log.Removed
	})

	return l
}

// Assume all arrays are of same length
func LogsAsArraysToLogs(arr LogsAsArrays) []*Log {
	numLogs := len(arr.Address)
	logs := make([]*Log, numLogs)

	slicefp.SliceEachCoroIdx(arr.BlockNumber, func(blockNumber int64, i int) {
		logs[i] = &Log{
			BlockNumber:      blockNumber,
			LogIndex:         arr.LogIndex[i],
			BlockHash:        arr.BlockHash[i],
			TransactionHash:  arr.TransactionHash[i],
			TransactionIndex: arr.TransactionIndex[i],
			Address:          arr.Address[i],
			Data:             arr.Data[i],
			Topic0:           arr.Topic0[i],
			Topic1:           arr.Topic1[i],
			Topic2:           arr.Topic2[i],
			Topic3:           arr.Topic3[i],
			Removed:          arr.Removed[i],
		}
	})

	return logs
}
