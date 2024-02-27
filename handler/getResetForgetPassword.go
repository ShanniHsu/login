package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetResetForgetPassword(ctx *gin.Context) {
	code := ctx.Query("code")
	email := ctx.Query("email")
	ctx.HTML(http.StatusOK, "resetForgetPassword.tmpl", gin.H{
		"code":  code,
		"email": email,
	})
}
