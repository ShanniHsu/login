package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetForgetPassword(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ForgetPassword.html", nil)
}
