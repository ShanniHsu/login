package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRegister(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "Register.html", nil)
}
