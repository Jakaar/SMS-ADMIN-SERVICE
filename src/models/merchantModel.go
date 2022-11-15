package model

import "github.com/google/uuid"

type Merchant struct {
	SelfID       uuid.UUID `json:"self_id"`
	Name         string    `json:"name"`
	UserID       string    `json:"user_id"`
	MerchantCode string    `json:"merchant_code"`
	Default
}

type MerchantName interface {
	TableName() string
}

func (Merchant) TableName() string {
	return "merchant"
}
