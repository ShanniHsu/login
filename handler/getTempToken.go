package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTempToken(ctx *gin.Context) {
	code := ctx.Query("code")
	ctx.HTML(http.StatusOK, "getTempToken.tmpl", gin.H{
		"code": code,
	})
}
