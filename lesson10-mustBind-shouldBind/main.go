package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Id string `form:"id" json:"id" xml:"id"  binding:"required,uuid"`
}

func main() {
	engine := gin.Default()
	engine.GET("user/", func(context *gin.Context) {
		var p Person
		//BindQuery 验证失败状态码为400
		//ShouldBindQuery 验证失败状态码为200
		err := context.ShouldBindQuery(&p)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"id": p.Id,
		})
	})
	engine.Run()
}
