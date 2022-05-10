package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/ping", ping)
	engine.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "pong",
		"data": "",
	})
}
