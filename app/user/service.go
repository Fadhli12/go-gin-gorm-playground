package user

import (
	"github.com/Fadhli12/go-gin-gorm-playground/app/auth"
	"github.com/Fadhli12/go-gin-gorm-playground/model"
)

type Service interface {
	FindAll() (model.Users, error)
	FindByID(ID int) (model.User, error)
	Create(userPost UserPost) (model.User, error)
	Update(ID int, userUpdate UserUpdate) (model.User, error)
	Delete(ID int) (model.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() (model.Users, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *service) FindByID(ID int) (model.User, error) {
	user, err := s.repository.FindByID(ID)
	return user, err
}

func (s *service) Create(userPost UserPost) (model.User, error) {
	user := model.User{
		Name:     userPost.Name,
		Email:    userPost.Email,
		Password: auth.HashAndSalt([]byte(userPost.Password)),
	}
	newUser, err := s.repository.Create(user)
	return newUser, err
}

func (s *service) Update(ID int, userUpdate UserUpdate) (model.User, error) {
	user, err := s.repository.FindByID(ID)

	user.Name = userUpdate.Name
	user.Email = userUpdate.Email
	if userUpdate.Password != "" {
		user.Password = auth.HashAndSalt([]byte(userUpdate.Password))
	}

	updateUser, err := s.repository.Update(user)
	return updateUser, err
}

func (s *service) Delete(ID int) (model.User, error) {
	user, err := s.repository.FindByID(ID)
	deletedUser, err := s.repository.Delete(user)
	return deletedUser, err
}
