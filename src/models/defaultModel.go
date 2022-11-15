package model

import (
	"gorm.io/gorm"
	"time"
)

type Default struct {
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
	UpdatedDate time.Time `json:"updated_date"`
	UpdatedBy   string    `json:"updated_by" gorm:"type:string;default:'admin'"`
	Status      bool      `gorm:"type:boolean;default:true" json:"status"`
}

func (d *Default) BeforeCreate(tx *gorm.DB) (err error) {
	d.Status = true
	d.CreatedDate = time.Now()
	d.UpdatedDate = time.Now()
	//d.CreatedBy = helper.DefaultFields(tx)
	return err
}
