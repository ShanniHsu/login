package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"login/router/repository/USER"
	"net/http"
	"strconv"
)

func WithdrawAmount(ctx *gin.Context) {
	token := ctx.PostForm("token")
	withdrawAmount := ctx.PostForm("withdraw_amount")

	user, err := USER.GetByToken(token)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "RespAmount.html", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    user.TempToken,
		})
		return
	}

	amount, err := strconv.ParseInt(withdrawAmount, 10, 64)
	if err != nil {
		err = errors.New("Please enter number!")
		ctx.HTML(http.StatusBadRequest, "RespAmount.html", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    user.TempToken,
		})
		return
	}

	amount = user.Amount - amount
	if amount < 0 {
		err = errors.New("Balance Amount isn't enough!")
		ctx.HTML(http.StatusBadRequest, "RespAmount.html", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    user.TempToken,
		})
		return
	}
	newData := map[string]interface{}{
		"amount": amount,
	}
	fmt.Println("amount: ", amount)

	err = USER.Update(user, newData)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "RespAmount.html", gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"code":    user.TempToken,
		})
		return
	}

	totalAmount := strconv.FormatInt(amount, 10)

	ctx.HTML(http.StatusOK, "RespAmount.html", gin.H{
		"status":       http.StatusOK,
		"message":      "Deposit Amount Successfully!",
		"code":         user.TempToken,
		"total_amount": "總金額:" + totalAmount,
	})
}
