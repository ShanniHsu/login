package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"login/pkg/jwt"
	"login/pkg/redis"
	"login/router/middleware"
	"login/router/repository/USER"
	"net/http"
)

func BuildTempToken(ctx *gin.Context) {
	userID := middleware.GetSession(ctx)
	user, err := USER.GetByID(userID)
	Ctx := context.TODO()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "RespTempToken.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	tempToken, err := jwt.GenerateToken()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "RespTempToken.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = redis.NewClient().Set(Ctx, "tempToken", tempToken, 0).Err()
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "RespTempToken.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	url := "http://127.0.0.1:8080/users-auth/token-change?code=" + tempToken

	newData := map[string]interface{}{
		"temp_token": tempToken,
	}
	err = USER.Update(user, newData)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "RespTempToken.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "RespTempToken.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "Build temp token successfully!!",
		"url":     url,
	})
}
