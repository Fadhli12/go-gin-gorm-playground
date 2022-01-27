package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookPost BookPost) (Book, error)
	Update(ID int, bookUpdate BookUpdate) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookPost BookPost) (Book, error) {
	price, _ := bookPost.Price.Int64()
	rating, _ := bookPost.Rating.Int64()
	discount, _ := bookPost.Discount.Int64()
	book := Book{
		Title:       bookPost.Title,
		Description: bookPost.Description,
		Price:       int(price),
		Rating:      int(rating),
		Discount:    int(discount),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookUpdate BookUpdate) (Book, error) {
	book, err := s.repository.FindByID(ID)

	price, _ := bookUpdate.Price.Int64()
	rating, _ := bookUpdate.Rating.Int64()
	discount, _ := bookUpdate.Discount.Int64()

	book.Title = bookUpdate.Title
	book.Description = bookUpdate.Description
	book.Price = int(price)
	book.Rating = int(rating)
	book.Discount = int(discount)

	updateBook, err := s.repository.Update(book)
	return updateBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	deletedBook, err := s.repository.Delete(book)
	return deletedBook, err
}
