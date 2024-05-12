package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"login/pkg/jwt"
	"login/pkg/rule"
	"login/router/middleware"
	"login/router/repository/USER"
	"net/http"
)

func Login(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	//var userLog = new(model.UserLoginLog)
	//檢查帳號
	user, err := USER.GetByAccount(account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The account isn't existed!")
		}
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"status":  http.StatusBadRequest,
		//	"message": err.Error(),
		//})

		ctx.HTML(http.StatusBadRequest, "ReturnLogin.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//檢查密碼
	if !rule.CheckPasswordHash(password, user.Password) {
		err = errors.New("The password isn't correct!")

		//紀錄Log
		//userLog.UserID = user.ID
		//userLog.Account = user.Account
		//userLog.Result = "0"
		//userLog.Remark = err.Error()
		//userLog.CreatedAt = time.Now()
		//userLog.UpdatedAt = time.Now()
		//
		//err = USER_LOGIN_LOG.Create(userLog)
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}

		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"status":  http.StatusBadRequest,
		//	"message": err.Error(),
		//})

		ctx.HTML(http.StatusBadRequest, "ReturnLogin.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//req := new(login_by_account.Request)
	//err := ctx.BindJSON(req)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, err)
	//	return
	//}
	//user := new(model.User)
	////檢查帳號
	//user, err = repository.GetByAccount(req.Account)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		err = errors.New("The account isn't existed!")
	//	}
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	////檢查密碼
	//if !ApplicationLogic.CheckPasswordHash(req.Password, user.Password) {
	//	err = errors.New("The password isn't correct!")
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}

	middleware.ClearSession(ctx)
	middleware.SaveSession(ctx, user.ID)

	token, err := jwt.GenerateToken()
	if err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"status":  http.StatusBadRequest,
		//	"message": err.Error(),
		//})

		ctx.HTML(http.StatusBadRequest, "ReturnLogin.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	newData := map[string]interface{}{
		"token": token,
	}

	err = USER.Update(user, newData)
	if err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"status":  http.StatusBadRequest,
		//	"message": err.Error(),
		//})

		ctx.HTML(http.StatusBadRequest, "ReturnLogin.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	//紀錄Log
	//userLog.UserID = user.ID
	//userLog.Account = user.Account
	//userLog.Result = "1"
	//userLog.CreatedAt = time.Now()
	//userLog.UpdatedAt = time.Now()
	//err = USER_LOGIN_LOG.Create(userLog)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	ctx.HTML(http.StatusOK, "LoginResp.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "You login successfully!",
	})
	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "You login successfully!",
	//	"data":    user,
	//})
}
