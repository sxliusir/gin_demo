package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

type User struct {
	ID        int    `db:"id"`
	Age       int    `db:"age"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

// 初始化数据库
func initDb() (err error) {
	db, err = sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/gin_demo")
	if err != nil {
		log.Fatalln(err)
		return err
	}
	db.SetMaxOpenConns(0)
	return nil
}

// 查询单条
func findOne() (err error) {
	var user User
	sqlStr := "SELECT * FROM users WHERE id=?"
	err = db.Get(&user, sqlStr, 22)
	if err != nil {
		fmt.Printf("数据查询失败")
		return err
	}
	fmt.Println(user)
	return nil
}

// 查询多条
func getMulti() {
	users := []User{}
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)
}

// 更新
func updateInfo() {
	exec, err := db.Exec("UPDATE users SET first_name='Bob Jonathan' WHERE id = 2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exec.RowsAffected())
}

// 插入数据
func insertData() {
	exec, err := db.Exec("INSERT INTO users (age, email, first_name, last_name) VALUES (?, ?, ?, ?)", 30,
		"tom@126.com", "tom", "liu")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exec.LastInsertId())
}

//删除数据
func del() {
	exec, err := db.Exec("DELETE FROM users WHERE id = 2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exec.RowsAffected())
}

func selectNamedQuery() {
	tmpU := User{Age: 18}
	//rows, err := db.NamedQuery(`SELECT * FROM users WHERE age=:age`, map[string]interface{}{"age": 18})
	rows, err := db.NamedQuery(`SELECT * FROM users WHERE age=:age`, tmpU)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("StructScan failed err:%v\n", err)
			continue
		}
		fmt.Println(u)
	}
}

//批量插入
func batchInsert() {
	userStructs := []User{
		{Age: 11, FirstName: "Savea", LastName: "Wang", Email: "wang@126.com"},
		{Age: 12, FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{Age: 13, FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}

	_, err := db.NamedExec(`INSERT INTO users (age, email, first_name, last_name) 
        VALUES (:age, :email, :first_name, :last_name)`, userStructs)
	if err != nil {
		fmt.Println("数据插入失败!")
	}
}
func main() {
	//初始化数据库
	err := initDb()
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功！")
	/*
		//查找单条
		err2 := findOne()
		if err2 != nil {
			fmt.Println(err2)
		}
		//查找多条
		getMulti()
		//更新数据
		updateInfo()
		//插入数据
		insertData()
	*/
	//删除数据
	//del()
	//selectNamedQuery()
	//批量插入
	batchInsert()
}
