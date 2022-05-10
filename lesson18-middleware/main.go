package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//中间件
func testMiddleware(context *gin.Context) {
	fmt.Println("中间件测试")
}
func userMiddleware(context *gin.Context) {
	fmt.Println("用户中间件测试")
}
func carMiddleware(context *gin.Context) {
	fmt.Println("汽车中间件测试")
}
func main() {
	engine := gin.Default()
	// 全局使用中间件
	engine.Use(testMiddleware)
	engine.GET("/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "login",
		})
	})
	engine.GET("/reg", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "reg",
		})
	})
	//组使用中间件
	group := engine.Group("/user", userMiddleware)
	{
		group.GET("/:id", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": context.Param("id"),
			})
		})
		group.POST("/:id", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": context.Param("id"),
			})
		})
	}
	//单个请求使用中间件
	engine.GET("/car/:id", carMiddleware, func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": context.Param("id"),
		})
	})
	engine.POST("/car/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": context.Param("id"),
		})
	})

	engine.Run()
}
