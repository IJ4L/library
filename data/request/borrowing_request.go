package request

type Borrowing struct {
	Id             int `json:"id" validate:"required" errormsg:"id is a required"`
	Borrowing_date int `json:"borrowing_date " validate:"required" errormsg:"borrowing_date is a required"`
	Return_date    int `json:"return_date" validate:"required" errormsg:"return_date is a required"`
	Book_id        int `json:"book_id" validate:"required" errormsg:"book_id is a required"`
	Member_id      int `json:"member_id" validate:"required" errormsg:"member_id is a required"`
}
