package controller

import (
	"github.com/gin-gonic/gin"
	"login/router/service"
)

type ApiController interface {
	Register(ctx *gin.Context)
	ForgetPassword(ctx *gin.Context)
	ResetForgetPassword(ctx *gin.Context)
	Login(ctx *gin.Context)
	BuildTempToken(ctx *gin.Context)
	TokenChange(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	DepositAmount(ctx *gin.Context)
	WithdrawAmount(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type apiController struct {
	userService service.UserService
}

func NewApiController(
	userService service.UserService,
) ApiController {
	return apiController{
		userService: userService,
	}
}
