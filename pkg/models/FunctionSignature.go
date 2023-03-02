package models

type FunctionSignature struct {
	HexSignature  string `gorm:"primaryKey"`
	TextSignature string
}

func (FunctionSignature) TableName() string {
	return "eth_function_signature"
}
