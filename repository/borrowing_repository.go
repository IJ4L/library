package repository

import (
	"library.com/data/response"
)

type BorrowingRepository interface {
	GetAllBorrowing() []response.Borrowings
}