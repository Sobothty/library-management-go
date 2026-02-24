package Repository

import (
	"Go-Project/feature/Model"

	"gorm.io/gorm"
)

type Gorm struct {
	db *gorm.DB
}

func (g Gorm) GetAllBook() ([]*Model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (g Gorm) GetBookById(id int) (*Model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (g Gorm) UpdateBook(book *Model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (g Gorm) CreateBook(book *Model.Book) error {
	//TODO implement me
	panic("implement me")
}

func (g Gorm) DeleteBook(id int) error {
	//TODO implement me
	panic("implement me")
}

func ConnectRepository(db *gorm.DB) BookRepository {
	return &Gorm{db: db}
}
