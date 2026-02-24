package Routes

import (
	"Go-Project/feature/Controller"
	"net/http"
)

type Routes struct {
	controller *Controller.BookController
}

func NewRoutes(Controller *Controller.BookController) *Routes {
	return &Routes{controller: Controller}
}

func (r *Routes) RoutesEndpoint(app *http.ServeMux) {
	//Book Endpoint
	app.HandleFunc("/books", r.controller.GetAllBook)
	app.HandleFunc("/book/{id}", r.controller.GetBookById)
	app.HandleFunc("/books/add", r.controller.AddBook)
	app.HandleFunc("/books/delete", r.controller.DeleteBook)
	app.HandleFunc("/books/borrow", r.controller.BorrowBook)
	app.HandleFunc("/books/return", r.controller.ReturnBook)
}
