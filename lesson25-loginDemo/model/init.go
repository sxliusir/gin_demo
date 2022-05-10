package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "log"
)

var DB *sqlx.DB

// 初始化数据库
func InitDb() (err error) {
	DB, err = sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/gin_demo")
	if err != nil {
		fmt.Println("connect 2 databse failed, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return nil
}
