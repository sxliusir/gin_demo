package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func middleware1(context *gin.Context) {
	log.Println("middleware1 in ...")
	context.Set("key", 1000)
	log.Println("middleware1 before next ...")
	k := context.GetInt("key")
	if k == 1000 {
		//context.Abort()
		context.AbortWithStatusJSON(http.StatusOK, gin.H{
			"msg": "验证失败",
		})
		return
	}
	context.Next()
	log.Println("middleware1 next after ....")
	log.Println("middleware1 done ....")

}

func middleware2(context *gin.Context) {
	log.Println("middleware2 in ...")
	log.Println("middleware2 before next ...")
	context.Next()
	log.Println("middleware2 next after ....")
	log.Println("middleware2 done ....")
}

func main() {
	engine := gin.Default()
	engine.Use(middleware1, middleware2)
	engine.GET("/ping", func(context *gin.Context) {
		log.Println("func in ...")
		k := context.GetInt("key")
		context.Set("key", k+2000)
		log.Println("func done...")
		context.JSON(http.StatusOK, gin.H{
			"msg": context.GetInt("key"),
		})
	})
	engine.Run()
}
