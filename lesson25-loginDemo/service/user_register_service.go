package service

import (
	"GinDemo/lesson25-loginDemo/model"
	"GinDemo/lesson25-loginDemo/serializer"
	"GinDemo/lesson25-loginDemo/utils"
)

type UserRegisterService struct {
	NickName string `json:"nickName" binding:"required,min=3,max=7" db:"nickName"`
	Age      int    `json:"age" binding:"required,gte=1,lte=150"`
	Gender   int    `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required,len=8"`
	Email    string `json:"email" binding:"required,email"`
}

func (service *UserRegisterService) Register() serializer.Response {
	sqlStr := `SELECT count(1) FROM USER WHERE email = ?`
	var count int
	_ = model.DB.Get(&count, sqlStr, service.Email)
	if count > 0 {
		return serializer.Response{
			Code:  40001,
			Data:  nil,
			Msg:   "邮箱已注册",
			Error: "",
		}
	}
	//密码加密
	service.Password = utils.Md5(service.Password)
	//创建用户
	sqlStr2 := `INSERT INTO user (nickName, age, gender, password, email) VALUES (:nickName, :age, :gender, :password, :email)`
	_, err := model.DB.NamedExec(sqlStr2, service)
	if err != nil {
		return serializer.Response{
			Code:  40002,
			Data:  nil,
			Msg:   "用户注册失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code:  200,
		Data:  nil,
		Msg:   "用户注册成功",
		Error: "",
	}
}
