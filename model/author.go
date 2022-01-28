package model

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name      string `json:"name" gorm:"type:varchar(100);not null"`
	Email     string `json:"email" gorm:"type:varchar(100);not null;unique;unique_index"`
	Biography string `json:"biography" gorm:"type:text;default:null"`
	Books     []Book
}

type Authors []Author
