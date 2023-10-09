package repository

import (
	"gorm.io/gorm"
	"library.com/data/response"
)

type BorrowingRepositoryImpl struct {
	DB *gorm.DB
}

func NewBorrowingRepositoryImpl(db *gorm.DB) *BorrowingRepositoryImpl {
	return &BorrowingRepositoryImpl{
		DB: db,
	}
}

func (repo *BorrowingRepositoryImpl) GetAllBorrowing() []response.Borrowings {
	var borrowings []response.Borrowings
	repo.DB.Preload("books").Preload("books").Preload("members").Find(&borrowings)
	return borrowings
}