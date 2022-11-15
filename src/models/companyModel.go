package model

import "github.com/google/uuid"

type Company struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	AvatarUrl      string    `json:"avatar_url"`
	Username       string    `json:"username"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	RegistryNumber string    `json:"registry_number"`
	UserType       string    `json:"user_type" gorm:"not null"`
	AreaOfActivity string    `json:"area_of_activity"`
	Province       string    `json:"province"`
	District       string    `json:"district"`
	Committee      string    `json:"committee"`
	Address        string    `json:"address"`
	Email          string    `json:"email" gorm:"unique;not null"`
	WalletId       string    `json:"wallet_id"`
	IsActive       bool      `json:"is_active"`
	IsVat          bool      `json:"is_vat"`
	Phone          int       `json:"phone" gorm:"unique;not null"`
	Password       string    `json:"password"`
	Default
}

type CompanyName interface {
	TableName() string
}

func (Company) TableName() string {
	return "user_company"
}
