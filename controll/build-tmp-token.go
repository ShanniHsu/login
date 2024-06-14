package controll

import (
	"github.com/gin-gonic/gin"
)

func BuildTmpToken(ctx *gin.Context) {
	//userID := middleware.GetSession(ctx)
	//user, err := USER.GetByID(userID)
	//Ctx := context.TODO()
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//}
	//
	//tempToken, err := jwt.GenerateToken()
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//err = redis.NewClient().Set(Ctx, "tempToken", tempToken, 0).Err()
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//url := "http://127.0.0.1:8080/users-auth/get-token?code=" + tempToken
	//
	//newData := map[string]interface{}{
	//	"temp_token": tempToken,
	//}
	//err = repository.Update(user, newData)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "Build temp token successfully!!",
	//	"url":     url,
	//	"code":    tempToken,
	//})
	return
}
