package author

import "database/sql"

type Service interface {
	FindAll() ([]Author, error)
	FindByID(ID int) (Author, error)
	Create(authorPost AuthorPost) (Author, error)
	Update(ID int, authorUpdate AuthorUpdate) (Author, error)
	Delete(ID int) (Author, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Author, error) {
	dataAuthors, err := s.repository.FindAll()
	return dataAuthors, err
}

func (s *service) FindByID(ID int) (Author, error) {
	dataAuthor, err := s.repository.FindByID(ID)
	return dataAuthor, err
}

func (s *service) Create(authorPost AuthorPost) (Author, error) {
	dataAuthor := Author{
		Name:      authorPost.Name,
		Email:     sql.NullString{authorPost.Email, true},
		Biography: sql.NullString{authorPost.Biography, true},
	}
	newAuthor, err := s.repository.Create(dataAuthor)
	return newAuthor, err
}

func (s *service) Update(ID int, authorUpdate AuthorUpdate) (Author, error) {
	dataAuthor, err := s.repository.FindByID(ID)
	dataAuthor.Name = authorUpdate.Name
	if authorUpdate.Email != nil {
		dataAuthor.Email = sql.NullString{*authorUpdate.Email, true}
	}
	if authorUpdate.Biography != nil {
		dataAuthor.Biography = sql.NullString{*authorUpdate.Biography, true}
	}
	updateAuthor, err := s.repository.Update(dataAuthor)
	return updateAuthor, err
}

func (s *service) Delete(ID int) (Author, error) {
	dataAuthor, err := s.repository.FindByID(ID)
	deletedAuthor, err := s.repository.Delete(dataAuthor)
	return deletedAuthor, err
}
