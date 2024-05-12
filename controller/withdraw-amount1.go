package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"login/router/repository/USER"
	"net/http"
	"strconv"
)

func WithdrawAmount1(ctx *gin.Context) {
	token := ctx.PostForm("token")
	withdrawAmount := ctx.PostForm("withdraw_amount")

	user, err := USER.GetByToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	amount, err := strconv.ParseInt(withdrawAmount, 10, 64)
	if err != nil {
		err = errors.New("Please enter number!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	amount = user.Amount - amount
	if amount < 0 {
		err = errors.New("Balance Amount isn't enough!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	newData := map[string]interface{}{
		"amount": amount,
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
		"status":       http.StatusOK,
		"message":      "Deposit Amount Successfully!",
		"total_amount": amount,
	})
}
