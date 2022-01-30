package book

import "encoding/json"

type BookPost struct {
	Title       string      `form:"title" json:"title" binding:"required"`
	Price       json.Number `form:"price" json:"price" binding:"required,number"`
	Description string      `form:"description" json:"description" binding:"required"`
	Rating      json.Number `form:"rating" json:"rating" binding:"required,number"`
	Discount    json.Number `form:"discount" json:"discount" binding:"required,number"`
	AuthorID    json.Number `form:"author_id" json:"author_id" binding:"required,number"`
}

type BookUpdate struct {
	Title       string      `form:"title" json:"title" binding:"required"`
	Price       json.Number `form:"price" json:"price" binding:"required,number"`
	Description string      `form:"description" json:"description" binding:"required"`
	Rating      json.Number `form:"rating" json:"rating" binding:"required,number"`
	Discount    json.Number `form:"discount" json:"discount" binding:"required,number"`
	AuthorID    json.Number `form:"author_id" json:"author_id" binding:"required,number"`
}
