package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"login/controll"
	"login/handler"
	"login/router/controller"
	"login/router/middleware"
	"login/router/repository"
	"login/router/service"
	"net/http"
)

func Router(r *gin.RouterGroup) {
	repo := repository.NewRepository()          //連線初始化 資料庫相關
	userService := service.NewUserService(repo) //User相關服務
	apiController := controller.NewApiController(userService)
	fmt.Println(repo)

	r.POST("/register", apiController.Register)              //註冊
	r.POST("/forget-password", apiController.ForgetPassword) //寄送忘記密碼

	//r.GET("/forget-password", controller.GetForgetPassword)
	r.POST("/reset-forget-password", apiController.ResetForgetPassword) //重設密碼

	//web
	r.GET("/reset-forget-password", handler.GetResetForgetPassword)
	r.GET("/login", handler.GetLogin)
	r.GET("/register", handler.GetRegister)
	r.GET("/forget-password", handler.GetForgetPassword)
	r.GET("/logout", handler.GetLogout)

	user := r.Group("/users-auth", middleware.SetSession())
	//user.GET("/test", test)
	user.POST("/login", apiController.Login) //登入

	user.Use(middleware.AuthSession())
	{
		user.POST("/build-temp-token", apiController.BuildTempToken) //建立臨時Token
		user.POST("/token-change", apiController.TokenChange)        //取真Token
		user.POST("/user-info", apiController.GetUserInfo)           //個人資訊
		user.POST("/deposit-amount", apiController.DepositAmount)    //存款
		user.POST("/withdraw-amount", apiController.WithdrawAmount)  //提款
		user.POST("/logout", apiController.Logout)                   //登出

		//純Postman使用
		user.POST("/build-tmp-token", controll.BuildTmpToken)    //建立臨時Token
		user.GET("/get-token", controll.GetToken)                //取永久Token
		user.POST("/get-user-info", controll.UserInfo)           //取個人資訊
		user.POST("/deposit-amount1", controll.DepositAmount1)   //存款
		user.POST("/withdraw-amount1", controll.WithdrawAmount1) //提款

		//web
		user.GET("/token-change", handler.GetTempToken)
	}

}

//純測試
func test(c *gin.Context) {
	str := []byte("ok")                      // 對於[]byte感到疑惑嗎？ 因為網頁傳輸沒有string的概念，都是要轉成byte字節方式進行傳輸
	c.Data(http.StatusOK, "text/plain", str) // 指定contentType為 text/plain，就是傳輸格式為純文字啦～
}
