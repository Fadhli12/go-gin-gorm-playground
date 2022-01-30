package book

import (
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (model.Books, error)
	FindByID(ID int) (model.Book, error)
	Create(book model.Book) (model.Book, error)
	Update(book model.Book) (model.Book, error)
	Delete(book model.Book) (model.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (model.Books, error) {
	var books model.Books
	err := r.db.Preload("Author").Find(&books).Error
	return books, err
}

func (r *repository) FindByID(ID int) (model.Book, error) {
	var book model.Book
	err := r.db.Preload("Author").First(&book, ID).Error
	return book, err
}

func (r *repository) Create(book model.Book) (model.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book model.Book) (model.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(book model.Book) (model.Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
