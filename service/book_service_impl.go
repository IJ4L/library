package service

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"library.com/data/request"
	"library.com/data/response"
	"library.com/repository"
	"library.com/utils"
)

type BookServiceImpl struct {
	bookRepository repository.BookRepository
	validate       *validator.Validate
}

func NewBookServiceImpl(bookRepository repository.BookRepository, validator *validator.Validate) *BookServiceImpl {
	return &BookServiceImpl{
		bookRepository: bookRepository,
		validate:       validator,
	}
}

func (service *BookServiceImpl) FetchAllBooks() []response.Book {
	return service.bookRepository.GetAllBooks()
}

func (service *BookServiceImpl) AddBook(book *request.Book) error {
	if err := utils.ValidateFunc(book, service.validate); err != nil {
		return err
	}

	if err := service.bookRepository.CreateBook(book); err != nil {
		return err
	}

	return nil
}

func (service *BookServiceImpl) GetBookDetailById(id *int) (response.BookDetailView, error) {
	if *id < 0 {
		return response.BookDetailView{}, fmt.Errorf("error: ID cannot be nil")
	}

	book, err := service.bookRepository.ViewBookDetail(id)
	if err != nil {
		return response.BookDetailView{}, err
	}

	return book, nil
}

func (service *BookServiceImpl) EditBook(book *request.Book) error {
	fmt.Println(book.Id)

	if err := utils.ValidateFunc(book, service.validate); err != nil {
		return err
	}

	if err := service.bookRepository.UpdateBook(book); err != nil {
		return err
	}

	return nil
}

func (service *BookServiceImpl) RemoveBook(id *int) error {
	if *id < 0 {
		return fmt.Errorf("error: ID cannot be nil")
	}

	if err := service.bookRepository.DeleteBook(id); err != nil {
		return err
	}

	return nil
}
