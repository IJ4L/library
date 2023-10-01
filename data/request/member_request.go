package request

type Member struct {
	Id              int    `json:"id"`
	Name            string `json:"name" validate:"required" errormgs:"name is a required"`
	Address         string `json:"address" validate:"required" errormgs:"address is a required"`
	IndentityNumber int `json:"identity_number" validate:"required" errormgs:"identiy_number is a required"`
}
