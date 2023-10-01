package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"library.com/data/request"
	"library.com/data/response"
	"library.com/service"
	"library.com/utils"
)

type BookController struct {
	bookService service.BookService
	ctx         *gin.Context
}

func NewBookController(bookService service.BookService, ctx *gin.Context) *BookController {
	return &BookController{
		bookService: bookService,
		ctx:         ctx,
	}
}

func (controller *BookController) GetBooks(ctx *gin.Context) {
	books := controller.bookService.FetchAllBooks()

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success get books",
		Data:    books,
	})
}

func (controller *BookController) AddBook(ctx *gin.Context) {
	book := request.Book{}
	ctx.ShouldBindJSON(&book)
	if err := controller.bookService.AddBook(&book); err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success add book",
		Data:    nil,
	})
}

func (controller *BookController) GetBookDetailById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	book, err := controller.bookService.GetBookDetailById(&id)
	if err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success get book detail",
		Data:    book,
	})
}

func (controller *BookController) EditBook(ctx *gin.Context) {
	book := request.Book{}
	ctx.ShouldBindJSON(&book)
	if err := controller.bookService.EditBook(&book); err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success edit book",
		Data:    nil,
	})
}

func (controller *BookController) RemoveBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}
	if err := controller.bookService.RemoveBook(&id); err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success remove book",
		Data:    nil,
	})
}
