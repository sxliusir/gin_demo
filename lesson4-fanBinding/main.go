package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//加载静态文件
	engine := gin.Default()
	//泛绑定
	/*
		请求 http://localhost:8080/posts/yezi/dead/info/detail
		返回结果
		action	"/dead/info/detail"
		name	"yezi"
	*/

	engine.GET("/posts/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})
	engine.Run()
}
