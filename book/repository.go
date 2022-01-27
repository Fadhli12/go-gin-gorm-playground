package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(dataBook Book) (Book, error)
	Update(dataBook Book) (Book, error)
	Delete(dataBook Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var dataBooks []Book
	err := r.db.Find(&dataBooks).Error
	return dataBooks, err
}

func (r *repository) FindByID(ID int) (Book, error) {
	var dataBook Book
	err := r.db.Find(&dataBook).Error
	return dataBook, err
}

func (r *repository) Create(dataBook Book) (Book, error) {
	err := r.db.Create(&dataBook).Error
	return dataBook, err
}

func (r *repository) Update(dataBook Book) (Book, error) {
	err := r.db.Save(&dataBook).Error
	return dataBook, err
}

func (r *repository) Delete(dataBook Book) (Book, error) {
	err := r.db.Delete(&dataBook).Error
	return dataBook, err
}
