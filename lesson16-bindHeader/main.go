package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Header struct {
	Referer string `header:"Referer" binding:"required"`
}

func main() {
	engine := gin.Default()
	engine.GET("user/", func(context *gin.Context) {
		var p Header
		err := context.ShouldBindHeader(&p)

		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"Referer": p.Referer,
		})
	})
	engine.Run()
}
