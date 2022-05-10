package main

import (
	"database/sql"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "log"
	_ "net/http"
)

/*
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);
INSERT INTO users (age, email, first_name, last_name)
VALUES (30, 'jon@calhoun.io', 'Jonathan', 'Calhoun');
INSERT INTO users (age, email, first_name, last_name)
VALUES (52, 'bob@smith.io', 'Bob', 'Smith');
INSERT INTO users (age, email, first_name, last_name)
VALUES (15, 'jerryjr123@gmail.com', 'Jerry', 'Seinfeld');
*/
var db *sql.DB

type User struct {
	ID        int    `json:"id"`
	Age       int    `json:"age"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// initDb 初始化数据库
func initDb() (err error) {
	db, err = sql.Open("mysql", "root:root@/gin_demo")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

/*
	getOne 获取单条数据
	参考资料： https://www.calhoun.io/querying-for-a-single-record-using-gos-database-sql-package/
*/
func getOne(id string) User {
	sqlStatement := "SELECT * FROM users WHERE id=?"
	var user User
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user
	case nil:
		fmt.Println(user)
	default:
		panic(err)
	}
	return user
}

// getMulti 查询多条数据
func getMulti() []User {
	sqlStr := "SELECT * FROM users"
	stmt, err := db.Prepare(sqlStr)
	//rows, err := db.Query(sqlStr)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	users := make([]User, 0)
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Age, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			log.Println(err)
			return nil
		}
		users = append(users, u)
	}
	return users
}

//update 更新数据
func update() {
	res, err := db.Exec("UPDATE users SET first_name='Jonathan' WHERE id = 1")
	if err != nil {
		fmt.Printf("更新失败！%v\n", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("未更新!")
	}
	if affected > 0 {
		fmt.Println("更新成功！")
	} else {
		fmt.Println("0条受影响！")
	}
}

// 删除数据
func del() {
	res, err := db.Exec("DELETE FROM users WHERE id = 1")
	if err != nil {
		fmt.Printf("删除失败！%v\n", err)
	}
	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("删除失败!")
	}
	if affected > 0 {
		fmt.Println("删除成功！")
	} else {
		fmt.Println("0条受影响！")
	}
}

// 插入数据
func insert() {
	sqlStr := "INSERT INTO users (age, email, first_name, last_name) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(sqlStr)
	res, err := stmt.Exec(19, "liurui@126.com", "meimei", "han")
	//res, err := db.Exec(sqlStr, 18, "sxliusir@126.com", "lei", "li")
	if err != nil {
		fmt.Println("插入数据失败！")
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Println("插入数据失败！")
		} else {
			fmt.Println("插入数据成功！id = ", id)
		}
	}
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Println("数据库连接失败！")
	} else {
		fmt.Println("数据库连接成功！")
	}
	/*
		engine := gin.Default()
		engine.GET("/user/:id", func(context *gin.Context) {
			id := context.Param("id")
			fmt.Println(id)
			user := getOne(id)
			context.JSON(http.StatusOK, gin.H{
				"info": user,
			})
		})
		engine.GET("/users", func(context *gin.Context) {
			multi := getMulti()
			context.JSON(http.StatusOK, gin.H{
				"data": multi,
			})
		})
		engine.Run()
	*/
	//更新操作
	//update()
	//删除操作
	//del()
	//插入数据
	insert()
}
