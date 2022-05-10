package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

type LoginForm struct {
	UserName   string `json:"userName" binding:"required,min=3,max=7"`
	Password   string `json:"password" binding:"required,len=8"`
	RePassword string `json:"rePassword" binding:"required,len=8"`
}

type RegForm struct {
	UserName string `json:"userName" binding:"required,min=3,max=7"`
	Password string `json:"password" binding:"required,len=8"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Sex      uint32 `json:"sex" binding:"required`
	Email    string `json:"email" binding:"required,email"`
}

//var trans ut.Translator

func main() {
	/*
		err := InitializeTrans()
		if err != nil {
			fmt.Println(err.Error())
		}*/
	if err := InitTrans("zh"); err != nil {
		fmt.Println("翻译器错误！")
		return
	}
	engine := gin.Default()
	engine.POST("/login", loginHandlers)
	engine.POST("/reg", regHandlers)
	engine.Run()
}

func regHandlers(context *gin.Context) {
	var l RegForm
	err := context.ShouldBindJSON(&l)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); !ok {
			// 如果错误不能转化，不能进行翻译，就返回错误信息
			context.JSON(http.StatusOK, gin.H{
				"error": errors.Error(),
			})
		} else {
			// 如果错误能进行翻译，就返回翻译后的错误信息
			context.JSON(http.StatusBadRequest, gin.H{
				"error": errors.Translate(trans),
			})
		}
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": l,
		"msg":  "登录成功",
	})
}

func loginHandlers(context *gin.Context) {
	var l LoginForm
	err := context.ShouldBindJSON(&l)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":  40001,
			"msg":   "登录失败，请检查用户名密码",
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
	})
}

/*
func InitializeTrans() (err error) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		zhT := zh.New()
		uni := ut.New(zhT, zhT)
		trans, _ := uni.GetTranslator("zh")

		err = zh_translations.RegisterDefaultTranslations(v, trans)
	}
	return
}*/

var trans ut.Translator

// 初始化一个翻译器函数
func InitTrans(locale string) (err error) {
	// 修改gin框架中的validator引擎属性，实现定制
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册一个获取json的tag自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	zhT := zh.New()              //中文翻译器
	enT := en.New()              //英文翻译器
	uni := ut.New(enT, zhT, enT) // 配置默认翻译器、以及可支持的翻译器
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) error", locale)
	}
	switch locale {
	case "en":
		_ = en_translations.RegisterDefaultTranslations(v, trans)
	case "zh":
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	default:
		_ = en_translations.RegisterDefaultTranslations(v, trans)
	}
	return

}
