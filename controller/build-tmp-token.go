package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"login/middleware"
	"login/redisC"
	"login/repository/USER"
	"login/service/ApplicationLogic"
	"net/http"
)

func BuildTmpToken(ctx *gin.Context) {
	userID := middleware.GetSession(ctx)
	user, err := USER.GetByID(userID)
	Ctx := context.TODO()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	tempToken, err := ApplicationLogic.GenerateToken()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = redisC.NewClient().Set(Ctx, "tempToken", tempToken, 0).Err()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	url := "http://127.0.0.1:8080/users-auth/get-token?code=" + tempToken

	newData := map[string]interface{}{
		"temp_token": tempToken,
	}
	err = USER.Update(user, newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Build temp token successfully!!",
		"url":     url,
		"code":    tempToken,
	})
}
