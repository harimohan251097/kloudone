package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type person struct {
	id   string `json:"id,omitempty"`
	name string `json:"name,omitempty"`
}

func main() {
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	var persons []person
	result, err := db.Query("SELECT * from person")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var a1 person
		err := result.Scan(&a1.id, &a1.name)
		if err != nil {
			panic(err.Error())
		}
		persons = append(persons, a1)
	}
	fmt.Print(persons)
}
