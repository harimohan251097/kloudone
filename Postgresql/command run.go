package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	//result, err := db.Query("create table product(Productid integer,Productname varchar(30),Producttype varchar(30),Productdesc varchar(100))")
	result, err := db.Query("insert into product values(1,'table','furniture','strong')")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	fmt.Print(result)
}
