package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"library.com/data/request"
	"library.com/data/response"
	"library.com/service"
	"library.com/utils"
)

type MemberController struct {
	memberService service.MemberService
	ctx           *gin.Context
}

func NewMemberController(memberService service.MemberService, ctx *gin.Context) *MemberController {
	return &MemberController{
		memberService: memberService,
		ctx:           ctx,
	}
}

func (controller *MemberController) GetMembers(ctx *gin.Context) {
	members := controller.memberService.FetchAllMembers()
	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success get members",
		Data:    members,
	})
}

func (controller *MemberController) GetMemberDetailById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	member, err := controller.memberService.GetMemberDetailById(&id)
	if err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success get member detail",
		Data:    member,
	})
}

func (controller *MemberController) AddMember(ctx *gin.Context) {
	member := request.Member{}
	ctx.ShouldBindJSON(&member)
	if err := controller.memberService.AddMember(&member); err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success add member",
		Data:    nil,
	})
}

func (controller *MemberController) UpdateMember(ctx *gin.Context) {
	member := request.Member{}
	ctx.ShouldBindJSON(&member)
	if err := controller.memberService.UpdateMember(&member); err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success update member",
		Data:    nil,
	})
}

func (controller *MemberController) DeleteMember(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	if err := controller.memberService.DeleteMember(&id); err != nil {
		utils.HandleError(ctx, 400, "Bad Request", err)
		return
	}

	ctx.JSON(200, response.Response{
		Status:  "OK",
		Message: "Success delete member",
		Data:    nil,
	})
}
