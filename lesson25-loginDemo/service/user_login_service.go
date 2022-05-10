package service

import (
	"GinDemo/lesson25-loginDemo/model"
	"GinDemo/lesson25-loginDemo/serializer"
	"GinDemo/lesson25-loginDemo/utils"
)

type UserLoginService struct {
	Password string `json:"password" binding:"required,len=8"`
	Email    string `json:"email" binding:"required,email"`
}

func (service *UserLoginService) Login() serializer.Response {
	//密码加密
	service.Password = utils.Md5(service.Password)
	sqlStr := `SELECT count(1) FROM USER WHERE email = ? and password = ?`
	var count int
	_ = model.DB.Get(&count, sqlStr, service.Email, service.Password)
	if count == 0 {
		return serializer.Response{
			Code:  40001,
			Data:  nil,
			Msg:   "用户名和密码错误",
			Error: "",
		}
	}
	return serializer.Response{
		Code:  200,
		Data:  nil,
		Msg:   "用户登录成功",
		Error: "",
	}
}
