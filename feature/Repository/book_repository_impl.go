package Repository

import (
	"Go-Project/feature/Model"

	"gorm.io/gorm"
)

type Gorm struct {
	db *gorm.DB
}

func (g Gorm) GetAllBook() ([]*Model.Book, error) {
	var books []*Model.Book
	result := g.db.Find(&books)
	return books, result.Error
}

func (g Gorm) GetBookById(id int) (*Model.Book, error) {
	var book *Model.Book
	err := g.db.First(&book, id).Error
	return book, err
}

func (g Gorm) UpdateBook(books *Model.Book, id int) string {
	var book *Model.Book
	err := g.db.First(&book, id).Error
	if err != nil {
		return "Book not found"
	}
	books.Title = book.Title
	books.IsBorrow = book.IsBorrow
	books.UpdateAt = book.UpdateAt
	g.db.Save(&book)

	return "Book updated successfully"

}

func (g Gorm) CreateBook(book *Model.Book) string {
	g.db.Create(book)
	return "Book created successfully"
}

func (g Gorm) DeleteBook(id int) string {
	var book *Model.Book
	g.db.Delete(book, id)
	return "Book deleted successfully"
}

func ConnectRepository(db *gorm.DB) BookRepository {
	return &Gorm{db: db}
}
