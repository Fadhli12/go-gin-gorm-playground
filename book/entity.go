package book

import (
	"github.com/Fadhli12/go-gin-gorm-playground/author"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string
	Description string
	Price       int
	Rating      int
	Discount    int
	AuthorID    uint
	Author      author.Author `gorm:"foreignKey:AuthorID"`
}
