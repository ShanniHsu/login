package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"login/repository/USER"
	"login/service/ApplicationLogic"
	"net/http"
)

func ResetForgetPassword(ctx *gin.Context) {
	code := ctx.PostForm("code")
	email := ctx.PostForm("email")
	user, err := USER.GetByToken(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The code isn't existed!")
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if user.Email != email {
		err = errors.New("The email is not yours!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	password := ctx.PostForm("password")
	hashPassword := ApplicationLogic.HashPassword(password)

	newData := map[string]interface{}{
		"password": hashPassword,
	}
	err = USER.Update(user, newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "ReturnLogin.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "You reset password successfully. Please back to login page!",
	})

	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "You reset password successfully. Please back to login page!",
	//})
	return
}
