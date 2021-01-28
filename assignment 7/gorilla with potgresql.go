package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Product struct {
	Productid   string `json:"Productid,omitempty"`
	Productname string `json:"Productname,omitempty"`
	Producttype string `json:"Producttype,omitempty"`
	Productdesc string `json:"Productdesc,omitempty"`
}

func GetProductEndpoint(w http.ResponseWriter, req *http.Request) {
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	var product []Product
	result, err := db.Query("SELECT * from product")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var a1 Product
		err := result.Scan(&a1.Productid, &a1.Productname, &a1.Producttype, &a1.Productdesc)
		if err != nil {
			panic(err.Error())
		}

		product = append(product, a1)
	}

	json.NewEncoder(w).Encode(product)
}

func GetProduct1Endpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	query := "SELECT * from product WHERE Productid =" + params["id"]
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var a1 Product
	for result.Next() {
		err := result.Scan(&a1.Productid, &a1.Productname, &a1.Producttype, &a1.Productdesc)
		if err != nil {
			panic(err.Error())
		}

	}
	json.NewEncoder(w).Encode(a1)
}

func CreateProductEndpoint(w http.ResponseWriter, req *http.Request) {
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	k := "insert into product values(" + keyVal["Productid"] + ",'" + keyVal["Productname"] + "','" + keyVal["Producttype"] + "','" + keyVal["Productdesc"] + "')"
	//{"Productname":"hari","Productid":"3","Producttype":"person","Productdesc":"Hari Mohan"}
	result, err := db.Query(k)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	fmt.Fprintf(w, "New post was created")

	var product []Product
	result1, err := db.Query("SELECT * from product")
	if err != nil {
		panic(err.Error())
	}
	defer result1.Close()

	for result1.Next() {
		var a1 Product
		err := result1.Scan(&a1.Productid, &a1.Productname, &a1.Producttype, &a1.Productdesc)
		if err != nil {
			panic(err.Error())
		}
		product = append(product, a1)
	}
	json.NewEncoder(w).Encode(product)
}

func DeleteProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	query := "delete from product where Productid=" + params["id"]
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	fmt.Fprintf(w, "post is deleted")

	var product []Product
	result1, err := db.Query("SELECT * from product")
	if err != nil {
		panic(err.Error())
	}
	defer result1.Close()

	for result1.Next() {
		var a1 Product
		err := result1.Scan(&a1.Productid, &a1.Productname, &a1.Producttype, &a1.Productdesc)
		if err != nil {
			panic(err.Error())
		}
		product = append(product, a1)
	}
	json.NewEncoder(w).Encode(product)
}

func PutProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	k := "update product set Productname='" + keyVal["Productname"] + "',Producttype='" + keyVal["Producttype"] + "',Productdesc='" + keyVal["Productdesc"] + "' where Productid=" + params["id"]
	//{"Productname":"hari123","Producttype":"person123","Productdesc":"Hari Mohan"}
	result, err := db.Query(k)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	fmt.Fprintf(w, " post was updated")

	var product []Product
	result1, err := db.Query("SELECT * from product")
	if err != nil {
		panic(err.Error())
	}
	defer result1.Close()

	for result1.Next() {
		var a1 Product
		err := result1.Scan(&a1.Productid, &a1.Productname, &a1.Producttype, &a1.Productdesc)
		if err != nil {
			panic(err.Error())
		}
		product = append(product, a1)
	}
	json.NewEncoder(w).Encode(product)
}

func main() {

	/*insert, err := db.Query("create table product(Productid integer,Productname varchar(30),Producttype varchar(20),Productdesc varchar(300))")
	insert, err := db.Query(" insert into product values(0,'machine','tech','machine learning')")
	insert, err := db.Query(" insert into product values(1,'welder','machnical','gas welder')")
	*/
	router := mux.NewRouter()
	router.HandleFunc("/product", GetProductEndpoint).Methods("GET")
	router.HandleFunc("/product/{id}", GetProduct1Endpoint).Methods("GET")
	router.HandleFunc("/product/{id}", PutProductEndpoint).Methods("PUT")
	router.HandleFunc("/product", CreateProductEndpoint).Methods("POST")
	router.HandleFunc("/product/{id}", DeleteProductEndpoint).Methods("DELETE")
	http.ListenAndServe(":123", router)
}
