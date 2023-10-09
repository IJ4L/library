package model

import (
	"time"

	"gorm.io/gorm"
	"library.com/data/response"
)

type Borrowings struct {
	gorm.Model
	BorrowingDate time.Time
	ReturnDate    time.Time
	BookID        uint
	MemberID      uint
	Book          response.Book   `gorm:"foreignKey:BookID"`
	Member        response.Member `gorm:"foreignKey:MemberID"`
}