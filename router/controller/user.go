package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	create_user "login/content/create-user"
	deposit_amount "login/content/deposit-amount"
	forget_password "login/content/forget-password"
	get_user_info "login/content/get-user-info"
	login_by_account "login/content/login-by-account"
	reset_forget_password "login/content/reset-forget-password"
	token_change "login/content/token-change"
	"net/http"
)

func (c apiController) Register(ctx *gin.Context) {
	var err error
	req := new(create_user.Request)

	//取註冊資料
	lastName := ctx.PostForm("last_name")
	firstName := ctx.PostForm("first_name")
	nickName := ctx.PostForm("nick_name")
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")
	email := ctx.PostForm("email")
	gender := ctx.PostForm("gender")

	if lastName == "" {
		err = errors.New("The last name is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if firstName == "" {
		err = errors.New("The first name is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if nickName == "" {
		err = errors.New("The nick name is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if account == "" {
		err = errors.New("The account is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if password == "" {
		err = errors.New("The password is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if email == "" {
		err = errors.New("The E-mail is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if gender == "" {
		err = errors.New("The gender is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.LastName = lastName
	req.FirstName = firstName
	req.NickName = nickName
	req.Account = account
	req.Password = password
	req.Email = email
	req.Gender = gender

	err = c.userService.Register(req)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "ReturnLogin.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "ReturnLogin.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "Register successfully!",
	})

	return
}

func (c apiController) ForgetPassword(ctx *gin.Context) {
	var err error
	req := new(forget_password.Request)
	email := ctx.PostForm("email")
	if email == "" {
		err = errors.New("The E-mail is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	req.Email = email

	err = c.userService.ForgetPassword(req)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "ReturnLogin.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "ReturnLogin.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "Send mail Successfully!",
	})
	return
}

func (c apiController) ResetForgetPassword(ctx *gin.Context) {
	var err error
	req := new(reset_forget_password.Request)
	code := ctx.PostForm("code")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	if code != "" {
		err = errors.New("The code is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if email == "" {
		err = errors.New("The E-mail is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if password == "" {
		err = errors.New("The password is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	req.Code = code
	req.Email = email
	req.Password = password
	err = c.userService.ResetForgetPassword(req)
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
	return
}

func (c apiController) Login(ctx *gin.Context) {
	var err error
	req := new(login_by_account.Request)
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	if account == "" {
		err = errors.New("The account is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if password == "" {
		err = errors.New("The password is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.Account = account
	req.Password = password

	err = c.userService.Login(ctx, req)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "ReturnLogin.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "LoginResp.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "You login successfully!",
	})
	return
}

func (c apiController) BuildTempToken(ctx *gin.Context) {
	url, err := c.userService.BuildTempToken(ctx)
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
	return
}

func (c apiController) TokenChange(ctx *gin.Context) {
	var err error
	req := new(token_change.Request)
	code := ctx.PostForm("code")

	if code == "" {
		err = errors.New("The code is required")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.Code = code

	token, err := c.userService.TokenChange(req)
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
		"token":   token,
	})
	return
}

func (c apiController) GetUserInfo(ctx *gin.Context) {
	var err error
	req := new(get_user_info.Request)
	token := ctx.PostForm("token")

	if token == "" {
		err = errors.New("The token is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.Token = token

	res, err := c.userService.GetUserInfo(req)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "RespUserInfo.tmpl", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "RespUserInfo.tmpl", gin.H{
		"status":  http.StatusOK,
		"message": "Get user info successfully!",
		"code":    res.TempToken,
		"user": gin.H{"id": res.ID,
			"last_name":  res.LastName,
			"first_name": res.FirstName,
			"nick_name":  res.NickName,
			"email":      res.Email,
			"gender":     res.Gender,
			"amount":     res.Amount,
		},
		//"user":    user,
	})

	return
}

func (c apiController) DepositAmount(ctx *gin.Context) {
	var err error
	req := new(deposit_amount.Request)
	token := ctx.PostForm("token")
	depositAmount := ctx.PostForm("deposit_amount")

	if token == "" {
		err = errors.New("The token is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if depositAmount == "" {
		err = errors.New("The deposit amount is required!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	req.Token = token
	req.DepositAmount = depositAmount

	res, err := c.userService.DepositAmount(req)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "RespAmount.html", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    res.TempToken,
		})
		return
	}

	ctx.HTML(http.StatusOK, "RespAmount.html", gin.H{
		"status":       http.StatusOK,
		"message":      "Deposit Amount Successfully!",
		"code":         res.TempToken,
		"total_amount": "總金額:" + res.TotalAmount,
	})
	return
}
