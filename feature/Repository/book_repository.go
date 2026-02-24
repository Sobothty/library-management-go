package Repository

import "Go-Project/feature/Model"

type BookRepository interface {
	GetAllBook() ([]*Model.Book, error)
	GetBookById(id int) (*Model.Book, error)
	UpdateBook(book *Model.Book) error
	CreateBook(book *Model.Book) error
	DeleteBook(id int) error
}
