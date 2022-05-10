package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.POST("/user/:id", userHandler)
	engine.Run()
}

func userHandler(context *gin.Context) {
	id := context.Param("id")
	userName := context.PostForm("userName")
	context.JSON(http.StatusOK, gin.H{
		"id":       id,
		"userName": userName,
	})
}
