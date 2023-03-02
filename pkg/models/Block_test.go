package models_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/vulcanlink/web3-models/pkg/models"
)

func TestMarshalBlock(t *testing.T) {
	baseFeePerGas := int64(0x657a9b02c)
	block := &models.Block{
		BaseFeePerGas: &baseFeePerGas,
		Difficulty:    "0x0",
		ExtraData:     common.Hex2Bytes("6275696c64657230783639"),
		GasLimit:      0x1c9c380,
		GasUsed:       0xf9011d,
		Hash:          common.HexToHash("0x1df2676f535bdf06f43e96c0e96f7ed9832525289dc9bf0a3a73160d4245bf07"),
		LogsBloom:     "0x5fe3e157684a14e3e88d35b6e334cff031b0ff47bfcefb44f5fba093fdfa5f686046b61d12614f6970feb24d5fe6f5ff5f1d9e1ebfae7cdeba6cf06b73bd3a15e9e5780b4547ddafbf17f5eec37cfce2cc6b3d8fe36798cb3cfb5fab9ad77fbdda538951f7eab5efbddd9b5f4c0f3cebca657fe2a5fb4e7fb150aade7e7b6edeb1f69aed23f7dc1c4fda5df7766bba6356aa96819bdb56eae724637e5e77dbbfbb881ff31bc37266f1777fecfae7e5bfffdbee6d21efdb7e987bae3695bfff7b7f3f2ad3782e16f687b3a2fff04faed6e7ebe3e6bdea92bd44f5fb7e1f7bb6e0ac5bfebf0aed7e3d52a68ffff7f1dcc6ee9f1b81b7dfeec54bbd5d7cfbc37f55",
		Miner:         common.HexToAddress("0x690b9a9e9aa1c9db991c7721a92d351db4fac990"),
		MixHash:       common.HexToHash("0x60010fe381358162ce40f036908005c6c4b8ac0c845e39df6a67a20bb16ebaed"),
		Nonce:         "0x0000000000000000",
		Number:        0xfeac15,
		ParentHash:    common.HexToHash("0x80d6eb4688e8e3e854cecd0a869703c14e4d899db9ab38c2fe3461f9772f3c10"),
		ReceiptsRoot:  common.HexToHash("0xbabd9562034f8eb677b16d74f83c5ed24d7a73404275a68ccab8ac5359b3bf09"),
		SHA3Uncles:    common.HexToHash("0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"),
		StateRoot:     common.HexToHash("0x01603db4e0ffe1374dc919aae4906e0358d879ab8008e9ecbb43acc58bf7d1af"),
		Timestamp:     0x63f73157,
		Transactions: []common.Hash{
			common.HexToHash("0xc85dcef6f7eeb8482720232f134bb8cf0964c260ccd5e7773fc198c0a484ecb1"),
			common.HexToHash("0x64e399de46535e68f12f5c8cbbfe78bd4d7f8f83582abc3006ffbfe50341c679"),
			common.HexToHash("0xbe3db73e46cc2dfcf4e054248a67029e7725f96a0c6aea5d043c3ad84b67992b"),
		},
		TransactionsRoot: common.HexToHash("0x26c7d2f744826b4c443635b65c098a5dafa57cc48454d3680afc71f5d6579691"),
		Uncles:           []common.Hash{},
	}

	jsonBlock, err := json.Marshal(block)
	assert.NoError(t, err)
	fmt.Println(string(jsonBlock))
}
