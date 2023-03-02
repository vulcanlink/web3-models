package models

type EventSignature struct {
	HexSignature  string `gorm:"primaryKey"`
	TextSignature string
}

func (EventSignature) TableName() string {
	return "eth_event_signature"
}
