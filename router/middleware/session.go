package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

const userkey = "shanni1234567890"

//Use cookie to store session id
func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userkey)) //設置生成sessionId的密鑰
	return sessions.Sessions("sessionid", store)
}

//User Auth Session Middle  User的中間件
func AuthSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionID := session.Get(userkey)
		fmt.Println("sessionID:", sessionID)
		if sessionID == nil {
			//Abort退出來 暫停
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message:": "此頁面需要登入",
			})
			return
		}
		ctx.Next()
	}
}

//Save Session for User
func SaveSession(ctx *gin.Context, userID int64) {
	session := sessions.Default(ctx)
	session.Set(userkey, userID)
	session.Save()
}

//Clear Session for User
func ClearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
}

//Get Session for User
func GetSession(ctx *gin.Context) int64 {
	session := sessions.Default(ctx)
	sessionID := session.Get(userkey)
	if sessionID == nil {
		return 0
	}
	return sessionID.(int64)
}

//Check Session for User
func CheckSession(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	sessionID := session.Get(userkey)
	//sessionID不等於空就回傳true
	return sessionID != nil
}
