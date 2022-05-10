package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Product struct {
	Code  string
	Price uint
	gorm.Model
}

type Student struct {
	gorm.Model
	UserName string `gorm:"column:user_name;varchar(255);index:idx_user_name;unique"`
	Age      uint
	Email    string
}

// TableName 指定表名，相对 NamingStrategy ，它的优先级高
//func (Student) TableName() string {
//	return "sys_stu"
//}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	dsn := "root:root@tcp(127.0.0.1:3306)/gin_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", //自定义表前缀
			SingularTable: true,   //表名使用单数
		},
	})
	if err != nil {
		panic("数据库连接失败")
	}
	// 迁移 schema
	//db.AutoMigrate(&Product{})
	// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	// Read
	//var product Product
	//db.First(&product, 1) // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 2300, "Code": "F421"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)

	db.AutoMigrate(&Student{})

	//插入数据
	/*
		stu := Student{UserName: "李雷"}
		res := db.Create(&stu)
		fmt.Println(res.RowsAffected)
		fmt.Println(res.Error)
	*/
	//批量插入
	/*
		var stus = []Student{{UserName: "韩梅梅"}, {UserName: "echo"}, {UserName: "kevin"}}
		db.Create(&stus)
		for _, v := range stus {
			fmt.Println(v.UserId)
		}*/

	//分批次插入
	/*
		var stus = []Student{{UserName: "李梅"}, {UserName: "tom"}, {UserName: "jack"}}
		db.CreateInBatches(&stus, 1)*/
	//查询
	/*
		var stu = Student{UserId: 6}
		result := db.First(&stu)
		// 检查 ErrRecordNotFound 错误
		is := errors.Is(result.Error, gorm.ErrRecordNotFound)
		if is {
			fmt.Println("没有相应的数据")
			return
		}
		fmt.Println(stu.UserId)
		fmt.Println(stu.UserName)

		//获取全部数据
		var stu1 []Student
		res := db.Find(&stu1)
		fmt.Printf("共 %d 条数据\n", res.RowsAffected)
		for _, v := range stu1 {
			fmt.Println(v.UserName)
		}
	*/

	//条件查询
	/*
		var stu Student
		db.Where("user_name = ?", "echo").First(&stu)
		fmt.Println(stu.UserName)
		fmt.Println(stu.UserId)
		// Slice of primary keys
		var stu1 []Student
		db.Where([]int64{1, 3, 7}).Find(&stu1)
		for _, v := range stu1 {
			fmt.Println(v.UserName)
		}*/

	//更新
	/*
		var stu Student
		db.First(&stu)

		stu.UserName = "jinzhu 2"
		db.Save(&stu)
	*/
	//条件更新
	//db.Model(&Student{}).Where("age = ?", 1).Update("email", "sxliusir@126.com")
	/*
		var stu = Student{
			UserId: 1,
		}
		db.Model(&stu).Update("user_name", "hello")
	*/

	//var users = []Student{{UserName: "jinzhu1"}, {UserName: "jinzhu2"}, {UserName: "jinzhu3"}}
	//db.Create(&users)

	//删除
	//软删除
	//var stu = Student{
	//	Model: gorm.Model{
	//		ID: 1,
	//	},
	//}
	//db.Delete(&stu)
	var stu []Student
	db.Find(&stu)
	for _, v := range stu {
		fmt.Println(v.ID, v.UserName)
	}
}
