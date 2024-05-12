package controller

import (
	"github.com/gin-gonic/gin"
	"login/router/middleware"
	"login/router/repository/USER"
	"net/http"
)

func Logout(ctx *gin.Context) {
	userID := middleware.GetSession(ctx)
	user, err := USER.GetByID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	newData := map[string]interface{}{
		"token": "",
	}

	err = USER.Update(user, newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	middleware.ClearSession(ctx)

	ctx.HTML(http.StatusOK, "ReturnLogin.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "User Sign out successfully!",
	})

	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "User Sign out successfully",
	//})
}
