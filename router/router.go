package router

import (
	"github.com/gin-gonic/gin"
	"library.com/controller"
)

func NewRouter(bookController *controller.BookController, memberController *controller.MemberController, borrowingController *controller.BorrowingController) *gin.Engine {
	service := gin.Default()
	router := service.Group("/api")

	{
		bookRouter := router.Group("/books")
		{
			bookRouter.GET("", bookController.GetBooks)
			bookRouter.POST("", bookController.AddBook)
			bookRouter.PUT("/:id", bookController.EditBook)
			bookRouter.DELETE("/:id", bookController.RemoveBook)
			bookRouter.GET("/:id", bookController.GetBookDetailById)
		}
	}
	
	{
		memberRouter := router.Group("/members")
		{
			memberRouter.GET("", memberController.GetMembers)
			memberRouter.POST("", memberController.AddMember)
			memberRouter.PUT("/:id", memberController.UpdateMember)
			memberRouter.DELETE("/:id", memberController.DeleteMember)
			memberRouter.GET("/:id", memberController.GetMemberDetailById)
		}
	}

	{
		borrowingRouter := router.Group("/borrowings")
		{
			borrowingRouter.GET("", borrowingController.Index)
		}
	}

	return service
}