package service

import (
	"github.com/go-playground/validator/v10"
	"library.com/data/response"
	"library.com/repository"
)

type BorrowingServiceImpl struct {
	borrowingRepository repository.BorrowingRepository
	validate            *validator.Validate
}

func NewBorrowingServiceImpl(borrowingRepository repository.BorrowingRepository, validate *validator.Validate) *BorrowingServiceImpl {
	return &BorrowingServiceImpl{
		borrowingRepository: borrowingRepository,
		validate:            validate,
	}
}

func (service *BorrowingServiceImpl) GetBorrowings() []response.Borrowings {
	return service.borrowingRepository.GetAllBorrowing()
}