package Controller

import (
	"Go-Project/feature/Service"
	"net/http"
)

type BookController struct {
	svc Service.BookService
}

func NewBookController(service Service.BookService) *BookController {
	return &BookController{svc: service}
}

func (controller *BookController) AddBook(write http.ResponseWriter, request *http.Request) {}

func (controller *BookController) GetAllBook(write http.ResponseWriter, request *http.Request) {}

func (controller *BookController) GetBookById(write http.ResponseWriter, request *http.Request) {}

func (controller *BookController) DeleteBook(write http.ResponseWriter, request *http.Request) {}

func (controller *BookController) BorrowBook(write http.ResponseWriter, request *http.Request) {}

func (controller *BookController) ReturnBook(write http.ResponseWriter, request *http.Request) {}
