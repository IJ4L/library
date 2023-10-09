package response

import "time"

type Borrowings struct {
	Id            int       `json:"id"`
	BorrowingDate time.Time `json:"borrowing_date"`
	ReturnDate    time.Time `json:"return_date"`
	BookID        uint      `json:"book_id"`
	MemberID      uint      `json:"member_id"`
	Book          Book      `json:"book" gorm:"foreignKey:BookID"`
	Member        Member    `json:"member" gorm:"foreignKey:MemberID"`
}