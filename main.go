package main

import (
	"github.com/gin-gonic/gin"
	"login/connect"
	"login/router"
)

func main() {

	r := gin.Default()
	// 提供靜態文件服務
	r.Static("/view/static", "./view/static")
	// 加載templates文件中所有的板模
	r.LoadHTMLGlob("view/templates/*")
	v1 := r.Group("")
	//Mysql Connect
	connect.GetDBConn()
	//Gin
	router.Router(v1)

	//Ctx := context.TODO()
	////Redis
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//fmt.Println("aaaaa")
	////存储一个键值
	//err := client.Set(Ctx, "myvalue", "123", 0).Err()
	//if err != nil {
	//	fmt.Println("set_err:", err)
	//}
	//fmt.Println("bbbb")
	////取一个键的值
	//value, err := client.Get(Ctx, "myvalue").Result()
	//if err == redis.Nil {
	//	fmt.Println("key does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("myvalue:", value)
	//}

	r.Run(":8080")
}
