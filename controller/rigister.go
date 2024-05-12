package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"login/model"
	"login/pkg/rule"
	"login/router/repository/USER"
	"net/http"
	"time"
)

func Register(ctx *gin.Context) {
	//取註冊資料
	lastName := ctx.PostForm("last_name")
	firstName := ctx.PostForm("first_name")
	nickName := ctx.PostForm("nick_name")
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	gender := ctx.PostForm("gender")

	//檢查此帳號是否存在

	user, err := USER.GetByAccount(account)
	if user.ID != 0 {
		err = errors.New("The account is existed!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//檢查此信箱是否存在
	user, err = USER.GetByEmail(email)
	if user.ID != 0 {
		err = errors.New("The e-mail is existed!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//model.User型態
	var data = new(model.User)
	data.LastName = lastName
	data.FirstName = firstName
	data.NickName = nickName
	data.Account = account
	data.Password = rule.HashPassword(password)
	data.Email = email
	data.Gender = gender
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	//Insert into Users
	err = USER.Create(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "ReturnLogin.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "Register successfully!",
	})

	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "Register successfully!",
	//})

	//確定註冊資料無誤
	//req := new(create_user.Request)
	//err := ctx.BindJSON(req)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, req)
	//	return
	//}
	//
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, "Parameter Error")
	//	return
	//}
	//user := new(model.User)
	////檢查此帳號是否存在
	//user, err = repository.GetByAccount(req.Account)
	//if user.ID != 0 {
	//	err = errors.New("The account is existed!")
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	////檢查此信箱是否存在
	//user, err = repository.GetByEmail(req.Email)
	//if user.ID != 0 {
	//	err = errors.New("The e-mail is existed!")
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//password := DomainLogic.HashPassword(req.Password)
	////model.User型態
	//var data = new(model.User)
	//data.LastName = req.LastName
	//data.FirstName = req.FirstName
	//data.NickName = req.NickName
	//data.Account = req.Account
	//data.Password = password
	//data.Email = req.Email
	//data.Gender = req.Gender
	//data.CreatedAt = time.Now()
	//data.UpdatedAt = time.Now()
	//
	////Insert into Users
	//err = repository.Create(data)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//ctx.JSON(http.StatusOK, "Register successfully!")
}
