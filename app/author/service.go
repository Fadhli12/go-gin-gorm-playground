package author

import (
	"github.com/Fadhli12/go-gin-gorm-playground/model"
)

type Service interface {
	FindAll() (model.Authors, error)
	FindByID(ID int) (model.Author, error)
	Create(authorPost AuthorPost) (model.Author, error)
	Update(ID int, authorUpdate AuthorUpdate) (model.Author, error)
	Delete(ID int) (model.Author, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() (model.Authors, error) {
	dataAuthors, err := s.repository.FindAll()
	return dataAuthors, err
}

func (s *service) FindByID(ID int) (model.Author, error) {
	dataAuthor, err := s.repository.FindByID(ID)
	return dataAuthor, err
}

func (s *service) Create(authorPost AuthorPost) (model.Author, error) {
	dataAuthor := model.Author{
		Name:      authorPost.Name,
		Email:     authorPost.Email,
		Biography: authorPost.Biography,
	}
	newAuthor, err := s.repository.Create(dataAuthor)
	return newAuthor, err
}

func (s *service) Update(ID int, authorUpdate AuthorUpdate) (model.Author, error) {
	dataAuthor, err := s.repository.FindByID(ID)
	dataAuthor.Name = authorUpdate.Name
	dataAuthor.Email = authorUpdate.Email
	if authorUpdate.Biography != "" {
		dataAuthor.Biography = authorUpdate.Biography
	}
	updateAuthor, err := s.repository.Update(dataAuthor)
	return updateAuthor, err
}

func (s *service) Delete(ID int) (model.Author, error) {
	dataAuthor, err := s.repository.FindByID(ID)
	deletedAuthor, err := s.repository.Delete(dataAuthor)
	return deletedAuthor, err
}
