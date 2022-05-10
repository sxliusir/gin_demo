package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Id string `json:"id" json:"id" xml:"id"  binding:"required,uuid"`
}

func main() {
	engine := gin.Default()
	engine.GET("user/", func(context *gin.Context) {
		var p Person
		err := context.BindJSON(&p)
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
