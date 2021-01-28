package main

import (
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Address   string `json:"address,omitempty"`
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("database connected")
	var people []Person
	result, err := db.Query("SELECT * from rest")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var p1 Person
		err := result.Scan(&p1.ID, &p1.Firstname, &p1.Lastname, &p1.Address)
		if err != nil {
			panic(err.Error())
		}
		people = append(people, p1)
	}
	json.NewEncoder(w).Encode(people)
}

/*
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	person.Firstname = params["name"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}
*/
func main() {

	/*insert, err := db.Query("create table rest(ID integer,Firstname varchar(10),Lastname varchar(10),Address varchar(30))")
	insert, err := db.Query("insert into rest values(0,'hari','mohan','bilaspur)")
	insert, err := db.Query("insert into rest values(1,'hari1','mohan1','bilaspur1')")
	insert, err := db.Query("insert into rest values(2,'hari2','mohan2','bilaspur2')")
	insert, err := db.Query("insert into rest values(3,'hari3','mohan3','bilaspur3')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	defer db.Close()*/
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPersonEndpoint).Methods("GET")
	/*router.HandleFunc("/people/{id}", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}/{name}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")*/
	http.ListenAndServe(":123", router)
}
