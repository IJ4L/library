package controller

import (
	"github.com/gin-gonic/gin"
	"library.com/data/response"
	"library.com/service"
)

type BorrowingController struct {
	borrowingService service.BorrowingService
	ctx              *gin.Context
}

func NewBorrowingControlerImpl(borrowingService service.BorrowingService, ctx *gin.Context) *BorrowingController {
	return &BorrowingController{
		borrowingService: borrowingService,
		ctx:              ctx,
	}
}

func (controller *BorrowingController) Index(ctx *gin.Context) {
	borrowings := controller.borrowingService.GetBorrowings()
	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success get borrowing",
		Data:    borrowings,
	})
}