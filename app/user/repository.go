package user

import (
	"github.com/Fadhli12/go-gin-gorm-playground/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (model.Users, error)
	FindByID(ID int) (model.User, error)
	Create(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(user model.User) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (model.Users, error) {
	var users model.Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindByID(ID int) (model.User, error) {
	var user model.User
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *repository) Create(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) Update(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) Delete(user model.User) (model.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
