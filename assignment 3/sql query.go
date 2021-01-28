package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("database connected")

	defer db.Close()

	insert, err := db.Query(" SELECT * from product")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(insert)
	defer insert.Close()

	/*insert, err := db.Query(" create table article(ID integer,Articlename varchar(20),type varchar(30),Articledesc varchar(100))")
	insert, err := db.Query(" insert into article values(0,'machine','tech','machine learning')")
	insert, err := db.Query(" insert into article values(1,'welding','machnical','gas welding')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	defer db.Close()*/
}
