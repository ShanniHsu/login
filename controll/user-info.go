package controll

import (
	"github.com/gin-gonic/gin"
)

func UserInfo(ctx *gin.Context) {
	//token := ctx.PostForm("token")
	//user, err := USER.GetByToken(token)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "Get user info successfully!",
	//	"code":    user.TempToken,
	//	"user": gin.H{"id": user.ID,
	//		"last_name":  user.LastName,
	//		"first_name": user.FirstName,
	//		"nick_name":  user.NickName,
	//		"email":      user.Email,
	//		"gender":     user.Gender,
	//		"amount":     user.Amount,
	//	},
	//})
	return
}
