package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//路由分组
	engine := gin.Default()
	api := engine.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			posts := v1.Group("/posts")
			{
				posts.GET("/", func(c *gin.Context) {
					c.String(http.StatusOK, "GET")
				})
				posts.POST("/", func(c *gin.Context) {
					c.String(http.StatusOK, "POST")
				})
				posts.PUT("/:id", func(c *gin.Context) {
					c.String(http.StatusOK, "PUT"+c.Param("id"))
				})
				posts.DELETE("/:id", func(c *gin.Context) {
					c.String(http.StatusOK, "DELETE"+c.Param("id"))
				})
			}
		}
	}
	engine.Run()
}
