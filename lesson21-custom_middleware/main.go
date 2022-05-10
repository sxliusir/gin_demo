package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Referer() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Referer")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "验证失败",
			})
			return
		}
		c.Next()
	}
}

func main() {
	engine := gin.Default()
	engine.Use(Referer())
	engine.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "验证成功",
		})
	})
	engine.Run()
}
