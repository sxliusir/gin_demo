package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	//GET
	engine.GET("/posts", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "sxliusir")
		lastname := context.Query("lastname")
		context.JSON(http.StatusOK, gin.H{
			"firstname": firstname,
			"lastname":  lastname,
		})
	})
	//QueryArray
	/*
		http://localhost:8080/array?ids=1,2,3,3,3
		{"ids":["1,2,3,3,3"]}
	*/
	engine.GET("/array", ArrayHandlers)
	//QueryMap
	/*
		http://localhost:8080/map?user[%22id%22]=1&user[%22name%22]=jack
		{
		  "data": {
		    "\"id\"": "1",
		    "\"name\"": "jack"
		  }
		}
	*/
	engine.GET("/map", MapHandlers)
	engine.Run()
}

func MapHandlers(context *gin.Context) {
	m := context.QueryMap("user")
	context.JSON(http.StatusOK, gin.H{
		"data": m,
	})
}

func ArrayHandlers(context *gin.Context) {
	ids := context.QueryArray("ids")
	context.JSON(http.StatusOK, gin.H{
		"ids": ids,
	})
}
