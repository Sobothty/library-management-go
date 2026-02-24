package Service

import (
	"Go-Project/feature/Model"
	"Go-Project/feature/Repository"
)

type bookServiceImpl struct {
	repo Repository.BookRepository
}

func (b bookServiceImpl) AddBook(title string) (*Model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookServiceImpl) DeleteBook(id int) (*Model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookServiceImpl) BorrowBook(id int) (*Model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookServiceImpl) ReturnBook(id int) (*Model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookServiceImpl) GetAllBook() ([]Model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookServiceImpl) GetBookById(id int) (Model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookServiceImpl(r Repository.BookRepository) BookService {
	return &bookServiceImpl{repo: r}
}
