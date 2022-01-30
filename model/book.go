package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `form:"title" json:"title" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:text;default:null"`
	Price       int    `json:"price" gorm:"type:int(10);not null;default:0"`
	Rating      int    `json:"rating" gorm:"type:int(1);not null;default:0"`
	Discount    int    `json:"discount" gorm:"type:int(10);not null;default:0"`
	AuthorID    uint   `json:"author_id" gorm:"not null"`
	Author      Author `gorm:"foreignKey:AuthorID"`
}

type Books []Book
