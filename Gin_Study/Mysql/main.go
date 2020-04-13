package main

/**
连接mysql，进行一系列操作
1. 连接数据库
2. 创建数据库表
3. 执行相关操作
 */

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
)

type User struct {
	Id int
	Name string
	Age int
}


func main()  {

	//设置连接参数
	connStr := "root:root@tcp(127.0.0.1:3306)/ginsql"

	//获取数据库实例
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
/*
	//创建数据库表
	//user表(id, name, age)
	_, err = db.Exec("create table user(id int auto_increment primary key, " +
									"name varchar(12) not null," +
									"age int default 1);")

	if err != nil {
		log.Fatal(err.Error())
		return
	}
*/

	//插入数据到数据库,?表示用可变参数代替
	_, err = db.Exec("insert into user(name, age) values(?, ?);", "youqu", 20)
	if err != nil {
		log.Fatal(err.Error())
		return
	} else {
		fmt.Println("Insert Success.")
	}

	//查询数据库
	rows, err := db.Query("select * from user;")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

SCAN:
	if rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(user.Id, "|", user.Name, "|", user.Age)
		goto SCAN
	}

}
