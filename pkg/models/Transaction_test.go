package models_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/vulcanlink/web3-models/pkg/models"
)

func TestMarshalTransactionNoLogs(t *testing.T) {
	to := common.HexToAddress("0x657e383edb9a7407e468acbcc9fe4c9730c7c275")

	transaction := models.Transaction{
		Hash:             common.HexToHash("0x961da0676d437c19e41afa9455670d5fcdc3916e119f133298b5d07cd971822b"),
		BlockHash:        common.HexToHash("0xe1d473f7e0c46ce0cc4b05e98a95201f9a36bbaa6e1199ab838eb48ef1e17083"),
		BlockNumber:      0xeafbed,
		To:               &to,
		TransactionIndex: 0xd7,
		From:             common.HexToAddress("0xfb4fd288fb13ee17a1c7d616bcf6fab5fc562285"),
		Gas:              0x44b80,
		GasPrice:         0x14236ee71,
		Nonce:            0xa2,
		Input:            common.Hex2Bytes("37295c5e00000000000000000000000000000000000000000000000000000000000003fe0000000000000000000000000000000000000000000000000000000000000000"),
		Value:            "0x11c37937e08000",
	}

	jsonTransaction, err := json.Marshal(transaction)
	assert.NoError(t, err)
	fmt.Println(string(jsonTransaction))
}

func TestMarshalTransaction(t *testing.T) {
	blockNumber := int64(0xeafbed)
	blockHash := common.HexToHash("0xe1d473f7e0c46ce0cc4b05e98a95201f9a36bbaa6e1199ab838eb48ef1e17083")
	to := common.HexToAddress("0x657e383edb9a7407e468acbcc9fe4c9730c7c275")
	transactionIndex := int64(0xd7)

	log := &models.Log{
		BlockNumber:      blockNumber,
		LogIndex:         0x1d7,
		BlockHash:        blockHash,
		TransactionHash:  common.HexToHash("0x961da0676d437c19e41afa9455670d5fcdc3916e119f133298b5d07cd971822b"),
		TransactionIndex: transactionIndex,
		Address:          common.HexToAddress("0x657e383edb9a7407e468acbcc9fe4c9730c7c275"),
		Data:             common.FromHex("0x00000000000000000000000000000000000000000000000000000000000003fe000000000000000000000000fb4fd288fb13ee17a1c7d616bcf6fab5fc5622850000000000000000000000000000000000000000000000000000000000000000"),
		Topic0:           common.HexToHash("0x8873f53f40d4865bac9c1e8998aef3351bb1ef3db1a6923ab09621cf1a6659a9"),
		Topic1:           common.HexToHash("0x000000000000000000000000fb4fd288fb13ee17a1c7d616bcf6fab5fc562285"),
		Topic2:           common.HexToHash("0x0000000000000000000000000000000000000000000000000000000063056a00"),
		Topic3:           common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
	}

	transaction := &models.Transaction{
		Hash:             common.HexToHash("0x961da0676d437c19e41afa9455670d5fcdc3916e119f133298b5d07cd971822b"),
		BlockHash:        common.HexToHash("0xe1d473f7e0c46ce0cc4b05e98a95201f9a36bbaa6e1199ab838eb48ef1e17083"),
		BlockNumber:      0xeafbed,
		To:               &to,
		TransactionIndex: 0xd7,
		From:             common.HexToAddress("0xfb4fd288fb13ee17a1c7d616bcf6fab5fc562285"),
		Gas:              0x44b80,
		GasPrice:         0x14236ee71,
		Nonce:            0xa2,
		Input:            common.Hex2Bytes("37295c5e00000000000000000000000000000000000000000000000000000000000003fe0000000000000000000000000000000000000000000000000000000000000000"),
		Value:            "0x11c37937e08000",
		Logs:             []*models.Log{log},
	}

	jsonTransaction, err := json.Marshal(transaction)
	assert.NoError(t, err)
	fmt.Println(string(jsonTransaction))
}
