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

func TokenChange(ctx *gin.Context) {
	code := ctx.PostForm("code")
	Ctx := context.TODO()
	value, err := redisC.NewClient().Get(Ctx, "tempToken").Result()
	fmt.Println("tempToken's value:", value)
	if err != nil {
		if err == redis.Nil {
			err = errors.New("The url is expired!")
		}
		ctx.HTML(http.StatusBadRequest, "RespTokenChange.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if value != code {
		err = errors.New("The url is error!")
		ctx.HTML(http.StatusBadRequest, "RespTokenChange.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := USER.GetByTempToken(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The code isn't existed!")
		}
		ctx.HTML(http.StatusBadRequest, "RespTokenChange.tmpl", gin.H{
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
		ctx.HTML(http.StatusBadRequest, "RespTokenChange.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "tokenChange.html", gin.H{
		"status":  http.StatusOK,
		"message": "Token change successfully!",
		"code":    code,
		"token":   user.Token,
	})
}
