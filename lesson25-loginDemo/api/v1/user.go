package v1

import (
	"GinDemo/lesson25-loginDemo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginHandler(context *gin.Context) {
	var s service.UserLoginService
	err := context.ShouldBindJSON(&s)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg": "验证失败",
		})
	} else {
		register := s.Login()
		context.JSON(http.StatusOK, register)
	}
}

func UserRegHandler(context *gin.Context) {
	var s service.UserRegisterService
	err := context.ShouldBindJSON(&s)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg": "验证失败",
		})
	} else {
		register := s.Register()
		context.JSON(http.StatusOK, register)
	}
}
