package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Id   string `form:"id" binding:"required,uuid"`
	Name string `form:"name" binding:"required,min=3,max=5"`
}

func main() {
	engine := gin.Default()
	engine.GET("user/", func(context *gin.Context) {
		var p Person
		err := context.ShouldBind(&p)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"id":   p.Id,
			"Name": p.Name,
		})
	})
	engine.Run()
}
