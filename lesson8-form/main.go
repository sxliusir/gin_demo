package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.POST("/user", userHandler)
	engine.POST("/form_array", formArrayHandler)
	engine.POST("/form_map", formMapHandler)
	engine.Run()
}

func formMapHandler(context *gin.Context) {
	user := context.PostFormMap("user")
	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func formArrayHandler(context *gin.Context) {
	ids := context.PostFormArray("ids")
	context.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}

func userHandler(context *gin.Context) {
	user := context.PostForm("user")
	age := context.DefaultPostForm("age", "18")
	context.JSON(http.StatusOK, gin.H{
		"user": user,
		"age":  age,
	})

}
