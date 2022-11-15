package model

import (
	"github.com/google/uuid"
	"time"
)

type Wallet struct {
	WalletId           uuid.UUID `gorm:"primaryKey, uniqueIndex" json:"wallet_id"`
	CurrentBalance     float64   `json:"current_balance"`
	LastDeposit        time.Time `json:"last_deposit"`
	PostpaidLimitation float64   `json:"postpaid_limitation"`
	PaymentSpending    float64   `json:"payment_spending"`
	IsActive           bool      `json:"is_active"`
	CreatedDate        time.Time `json:"created_date"`
	CreatedBy          string    `json:"created_by"`
	UpdatedDate        time.Time `json:"updated_date"`
	UpdatedBy          string    `json:"updated_by" gorm:"type:string;default:'admin'"`
	Status             bool      `gorm:"type:boolean;default:true" json:"status"`
}

type WalletName interface {
	TableName() string
}

func (Wallet) TableName() string {
	return "wallet"
}
