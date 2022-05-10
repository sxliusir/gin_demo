package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//加载静态文件
	engine := gin.Default()
	//GET
	engine.Static("/images", "./images")
	//StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
	engine.StaticFS("/static", http.Dir("./static"))
	//加载单独的静态文件
	engine.StaticFile("index", "index.html")
	engine.Run()
}
