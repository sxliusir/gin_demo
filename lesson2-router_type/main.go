package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//多种请求类型
	engine := gin.Default()
	//GET
	engine.GET("/posts", func(context *gin.Context) {
		context.String(http.StatusOK, "GET")
	})
	//POST
	engine.POST("/posts", func(context *gin.Context) {
		context.String(http.StatusOK, "POST")
	})
	//PUT
	engine.PUT("/posts/:id", func(context *gin.Context) {
		context.String(http.StatusOK, context.Param("id"))
	})
	//DELETE
	engine.DELETE("/posts/:id", func(context *gin.Context) {
		context.String(http.StatusOK, "DELETE")
	})
	//Any
	engine.Any("/infos", func(context *gin.Context) {
		context.String(http.StatusOK, "Any")
	})
	engine.Run()
}
