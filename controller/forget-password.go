package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"login/repository/USER"
	"login/service/ApplicationLogic"
	"net/http"
	"strings"
)

func ForgetPassword(ctx *gin.Context) {
	email := ctx.PostForm("email")
	//檢查Email是否存在
	user, err := USER.GetByEmail(email)
	fmt.Println(user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The email isn't existed!")
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	var token string
	token, err = ApplicationLogic.GenerateToken()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	fmt.Println("忘記密碼:", token)
	newData := map[string]interface{}{
		"token": token,
	}
	//存token至user資料中
	err = USER.Update(user, newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	emailReplace := strings.Replace(email, "@", "%40", -1)
	url := "http://127.0.0.1:8080/reset-forget-password?code=" + token + "&" + "email=" + emailReplace
	subject := "ShanniTest- 請重置密碼!"
	body := "請點擊此連結: " + url + " ,並重新設置密碼。"

	//寄信
	err = ApplicationLogic.SendMail(subject, email, body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "ReturnLogin.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "Send mail Successfully!",
	})

	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "Send mail Successfully!",
	//})
	//req := new(forget_password.Request)
	//err := ctx.BindJSON(req)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, err)
	//	return
	//}
	////檢查帳號
	//user := new(model.User)
	//user, err = repository.GetByAccount(req.Account)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		err = errors.New("The account isn't existed!")
	//	}
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	////檢查Email是否存在
	//if user.Email != req.Email {
	//	err = errors.New("The email isn't existed!")
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}

}
