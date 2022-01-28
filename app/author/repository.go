package author

import (
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (model.Authors, error)
	FindByID(ID int) (model.Author, error)
	Create(author model.Author) (model.Author, error)
	Update(author model.Author) (model.Author, error)
	Delete(author model.Author) (model.Author, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (model.Authors, error) {
	var authors model.Authors
	err := r.db.Find(&authors).Error
	return authors, err
}

func (r *repository) FindByID(ID int) (model.Author, error) {
	var author model.Author
	err := r.db.First(&author, ID).Error
	return author, err
}

func (r *repository) Create(author model.Author) (model.Author, error) {
	err := r.db.Create(&author).Error
	return author, err
}

func (r *repository) Update(author model.Author) (model.Author, error) {
	err := r.db.Save(&author).Error
	return author, err
}

func (r *repository) Delete(author model.Author) (model.Author, error) {
	err := r.db.Delete(&author).Error
	return author, err
}
