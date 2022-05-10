package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Id   int    `uri:"id"`
	Name string `uri:"name"`
}

func main() {
	engine := gin.Default()
	//GET
	engine.GET("/:id/:name", func(c *gin.Context) {
		var p Person
		err := c.ShouldBindUri(&p)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":   p.Id,
			"name": p.Name,
		})
	})
	engine.Run()
}
