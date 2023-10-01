package service

import (
	"library.com/data/request"
	"library.com/data/response"
)

type BookService interface {
	FetchAllBooks() []response.Book
	AddBook(book *request.Book) error
	GetBookDetailById(id *int) (response.BookDetailView, error)
	EditBook(book *request.Book) error
	RemoveBook(id *int) error
}