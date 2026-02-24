package Service

import "Go-Project/feature/Model"

type BookService interface {
	AddBook(title string) (*Model.Book, error)
	DeleteBook(id int) (*Model.Book, error)
	BorrowBook(id int) (*Model.Book, error)
	ReturnBook(id int) (*Model.Book, error)
	GetAllBook() ([]Model.Book, error)
	GetBookById(id int) (Model.Book, error)
}
