package author

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Author, error)
	FindByID(ID int) (Author, error)
	Create(dataAuthor Author) (Author, error)
	Update(dataAuthor Author) (Author, error)
	Delete(dataAuthor Author) (Author, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Author, error) {
	var dataAuthors []Author
	err := r.db.Find(&dataAuthors).Error
	return dataAuthors, err
}

func (r *repository) FindByID(ID int) (Author, error) {
	var dataAuthor Author
	err := r.db.Find(&dataAuthor).Error
	return dataAuthor, err
}

func (r *repository) Create(dataAuthor Author) (Author, error) {
	err := r.db.Create(&dataAuthor).Error
	return dataAuthor, err
}

func (r *repository) Update(dataAuthor Author) (Author, error) {
	err := r.db.Save(&dataAuthor).Error
	return dataAuthor, err
}

func (r *repository) Delete(dataAuthor Author) (Author, error) {
	err := r.db.Delete(&dataAuthor).Error
	return dataAuthor, err
}
