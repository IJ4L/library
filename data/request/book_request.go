package request

type Book struct {
	Id          int    `json:"id"`
	Isbn        string `json:"isbn" validate:"required" errormgs:"ISBN is a required"`
	Title       string `json:"title" validate:"required" errormgs:"Title is a required"`
	Stock       int    `json:"stock" validate:"required" errormgs:"Stock is a required"`
	PublisherId int    `json:"publisher" validate:"required" errormgs:"Publisher is a required"`
}
