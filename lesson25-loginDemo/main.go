package main

import (
	v1 "GinDemo/lesson25-loginDemo/api/v1"
	"GinDemo/lesson25-loginDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	err := model.InitDb()
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	v := engine.Group("api/v1")
	{
		// ping
		v.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg" : "pong",
			})
		})
		// 用户
		userGroup := v.Group("/user")
		{
			// 注册
			userGroup.POST("/reg", v1.UserRegHandler)
			// 登录
			userGroup.POST("/login", v1.UserLoginHandler)
		}
	}
	engine.Run()
}
