package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "Login.html", nil)
}
