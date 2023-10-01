package repository

import (
	"library.com/data/request"
	"library.com/data/response"
)

type BookRepository interface {
	GetAllBooks() []response.Book
	CreateBook(*request.Book) error
	ViewBookDetail(id *int) (response.BookDetailView, error)
	UpdateBook(*request.Book) error
	DeleteBook(id *int) error
}
