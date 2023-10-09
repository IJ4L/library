package repository

import (
	"fmt"

	"gorm.io/gorm"
	"library.com/data/request"
	"library.com/data/response"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepositoryImpl(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{
		DB: db,
	}
}

func (repo *BookRepositoryImpl) GetAllBooks() []response.Book {
	var books []response.Book
	repo.DB.Find(&books)
	return books
}

func (repo *BookRepositoryImpl) CreateBook(book *request.Book) error {
	db := repo.DB.Table("books").Create(&book)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (repo *BookRepositoryImpl) ViewBookDetail(id *int) (response.BookDetailView, error) {
	var book response.BookDetailView
	db := repo.DB.Table("books").Preload("Publisher").Where("id = ?", id).First(&book)
	if db.Error != nil {
		return response.BookDetailView{}, db.Error
	}
	return book, nil
}

func (repo *BookRepositoryImpl) UpdateBook(book *request.Book) error {
	fmt.Println(book.Id)
	db := repo.DB.Table("books").Where("id = ?", book.Id).Updates(&book)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (repo *BookRepositoryImpl) DeleteBook(id *int) error {
	var book response.Book
	db := repo.DB.Table("books").Where("id = ?", &id).Delete(&book)
	if db.Error != nil {
		return db.Error
	}
	return nil
}