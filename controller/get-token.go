package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"login/redisC"
	"login/repository/USER"
	"net/http"
)

func GetToken(ctx *gin.Context) {
	code := ctx.Query("code")
	Ctx := context.TODO()
	value, err := redisC.NewClient().Get(Ctx, "tempToken").Result()

	if err != nil {
		if err == redis.Nil {
			err = errors.New("The radis is error!")
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	fmt.Println("value:", value)
	fmt.Println("code:", code)
	if value != code {
		err = errors.New("The url is error!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := USER.GetByTempToken(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The url is invalid!")
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	newData := map[string]interface{}{
		"temp_token": "",
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
		"message": "Get Token successfully!",
		"token":   user.Token,
	})
	return
}
