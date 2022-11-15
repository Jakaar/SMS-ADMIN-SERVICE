package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID    string
	Title string
	Body  string
}
