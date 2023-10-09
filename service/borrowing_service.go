package service

import "library.com/data/response"

type BorrowingService interface {
	GetBorrowings() []response.Borrowings
}
