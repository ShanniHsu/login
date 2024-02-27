package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLogout(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "Logout.html", nil)
}
