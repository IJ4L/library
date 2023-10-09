package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"library.com/config"
	"library.com/controller"
	"library.com/repository"
	"library.com/router"
	"library.com/service"
)

func main() {
	db := config.InitDatabase()
	validate := validator.New()
	var ctx *gin.Context

	bookRepository := repository.NewBookRepositoryImpl(db)
	bookService := service.NewBookServiceImpl(bookRepository, validate)
	bookController := controller.NewBookController(bookService, ctx)

	memberRepository := repository.NewMemberRepositoryImpl(db)
	memberService := service.NewMemberServiceImpl(memberRepository, validate)
	memberController := controller.NewMemberController(memberService, ctx)

	borrowingRepository := repository.NewBorrowingRepositoryImpl(db)
	borrowingService := service.NewBorrowingServiceImpl(borrowingRepository, validate)
	borrowingController := controller.NewBorrowingControlerImpl(borrowingService, ctx)

	routes := router.NewRouter(bookController, memberController, borrowingController)

	server := &http.Server{
		Addr:           "localhost:3000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
	// helper.ErrorPanic(server_err)
}