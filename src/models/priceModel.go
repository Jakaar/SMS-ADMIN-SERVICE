package model

type Price struct {
	ID      string `gorm:"primaryKey" json:"user_id"`
	Unitel  int    `json:"unitel"`
	Mobicom int    `json:"mobicom"`
	Skytel  int    `json:"skytel"`
	Gmobile int    `json:"gmobile"`
	Default
}

type PriceName interface {
	TableName() string
}

func (Price) TableName() string {
	return "price"
}
