package Repository

import "Go-Project/feature/Model"

type BookRepository interface {
	GetAllBook() ([]*Model.Book, error)
	GetBookById(id int) (*Model.Book, error)
	UpdateBook(book *Model.Book, id int) string
	CreateBook(book *Model.Book) string
	DeleteBook(id int) string
}
