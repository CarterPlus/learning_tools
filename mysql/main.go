package main

import (
	"database/sql"
	"fmt"

	//导入mysql驱动包"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	var user []User
	rows, err := db.Query("select * from user")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var Id int
		var Name string

		if err := rows.Scan(&Id, &Name); err != nil {
			continue
		}
		user = append(user, User{Id: Id, Name: Name})
	}

	fmt.Println(user)
	// [{1 红红的太阳在东边升起} {2 迷人的小花蕊}]
}
