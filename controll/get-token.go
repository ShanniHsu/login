package controll

import (
	"github.com/gin-gonic/gin"
)

func GetToken(ctx *gin.Context) {
	//code := ctx.Query("code")
	//Ctx := context.TODO()
	//value, err := redis.NewClient().Get(Ctx, "tempToken").Result()
	//
	//if err != nil {
	//	if err == r1.Nil {
	//		err = errors.New("The radis is error!")
	//	}
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//fmt.Println("value:", value)
	//fmt.Println("code:", code)
	//if value != code {
	//	err = errors.New("The url is error!")
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//user, err := USER.GetByTempToken(code)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		err = errors.New("The url is invalid!")
	//	}
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status":  http.StatusBadRequest,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//newData := map[string]interface{}{
	//	"temp_token": "",
	//}
	//
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
	//	"message": "Get Token successfully!",
	//	"token":   user.Token,
	//})
	return
}
