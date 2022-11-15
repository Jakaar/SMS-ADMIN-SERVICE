package model

type Admin struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Username  string `json:"username,omitempty"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Default
}

type AdminName interface {
	TableName() string
}

func (Admin) TableName() string {
	return "admin_user"
}
