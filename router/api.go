package router

import (
	"github.com/gin-gonic/gin"
	"login/controller"
	"login/handler"
	"login/middleware"
	"net/http"
)

func Router(r *gin.RouterGroup) {
	r.POST("/register", controller.Register)
	r.POST("/forget-password", controller.ForgetPassword)

	//r.GET("/forget-password", controller.GetForgetPassword)
	r.POST("/reset-forget-password", controller.ResetForgetPassword)

	//web
	r.GET("/reset-forget-password", handler.GetResetForgetPassword)
	r.GET("/login", handler.GetLogin)
	r.GET("/register", handler.GetRegister)
	r.GET("/forget-password", handler.GetForgetPassword)
	r.GET("/logout", handler.GetLogout)

	user := r.Group("/users-auth", middleware.SetSession())
	//user.GET("/test", test)
	user.POST("/login", controller.Login)

	user.Use(middleware.AuthSession())
	{
		user.POST("/build-temp-token", controller.BuildTempToken) //建立臨時Token
		user.POST("/token-change", controller.TokenChange)        //取真Token
		user.POST("/user-info", controller.GetUserInfo)           //個人資訊
		user.POST("/deposit-amount", controller.DepositAmount)    //存款
		user.POST("/withdraw-amount", controller.WithdrawAmount)  //提款
		user.POST("/logout", controller.Logout)                   //登出

		//純Postman使用
		user.POST("/build-tmp-token", controller.BuildTmpToken)    //建立臨時Token
		user.GET("/get-token", controller.GetToken)                //取永久Token
		user.POST("/get-user-info", controller.UserInfo)           //取個人資訊
		user.POST("/deposit-amount1", controller.DepositAmount1)   //存款
		user.POST("/withdraw-amount1", controller.WithdrawAmount1) //提款

		//web
		user.GET("/token-change", handler.GetTempToken)
	}

}

//純測試
func test(c *gin.Context) {
	str := []byte("ok")                      // 對於[]byte感到疑惑嗎？ 因為網頁傳輸沒有string的概念，都是要轉成byte字節方式進行傳輸
	c.Data(http.StatusOK, "text/plain", str) // 指定contentType為 text/plain，就是傳輸格式為純文字啦～
}
