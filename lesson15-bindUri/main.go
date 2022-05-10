package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func main() {
	engine := gin.Default()
	//http://localhost:8080/user/e853fde3-8ba3-4ddb-bca7-984d44598414
	engine.GET("user/:id", func(context *gin.Context) {
		var p Person
		err := context.BindUri(&p)

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
