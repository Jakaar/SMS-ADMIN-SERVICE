package model

import uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

type User struct {
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

type UserName interface {
	TableName() string
}

func (User) TableName() string {
	return "user"
}

//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	u.UUID = uuid.New()
//	return
//}
