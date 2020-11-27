package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Product struct {
	productid   string `json:"productid,omitempty"`
	productname string `json:"productname,omitempty"`
	Producttype string `json:"Producttype,omitempty"`
	productdesc string `json:"productdesc,omitempty"`
}

func GetProductEndpoint(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	var product []Product
	result, err := db.Query("SELECT * from product")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var a1 Product
		err := result.Scan(&a1.productid, &a1.productname, &a1.Producttype, &a1.productdesc)
		if err != nil {
			panic(err.Error())
		}

		product = append(product, a1)
	}
	fmt.Print(product)
	json.NewEncoder(w).Encode(product)
}

func GetProduct1Endpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("SELECT * from product WHERE productid = ?", params["productid"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var a1 Product
	for result.Next() {
		err := result.Scan(&a1.productid, &a1.productname, &a1.Producttype, &a1.productdesc)
		if err != nil {
			panic(err.Error())
		}

	}
	json.NewEncoder(w).Encode(a1)
}

func CreateProductEndpoint(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	k := "insert into product values(" + keyVal["productid"] + ",'" + keyVal["productname"] + "','" + keyVal["Producttype"] + "','" + keyVal["productdesc"] + "')"
	//{"productname":"hari","productid":"3","Producttype":"person","productdesc":"Hari Mohan"}
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
		err := result1.Scan(&a1.productid, &a1.productname, &a1.Producttype, &a1.productdesc)
		if err != nil {
			panic(err.Error())
		}
		product = append(product, a1)
	}
	json.NewEncoder(w).Encode(product)
}

func DeleteProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("delete from product where productid=?", params["id"])
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
		err := result1.Scan(&a1.productid, &a1.productname, &a1.Producttype, &a1.productdesc)
		if err != nil {
			panic(err.Error())
		}
		product = append(product, a1)
	}
	json.NewEncoder(w).Encode(product)
}

func PutProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	k := "update product set productname='" + keyVal["productname"] + "',type='" + keyVal["Producttype"] + "',productdesc='" + keyVal["productdesc"] + "' where productid=" + params["id"]
	//{"productname":"hari123","Producttype":"person123","productdesc":"Hari Mohan"}
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
		err := result1.Scan(&a1.productid, &a1.productname, &a1.Producttype, &a1.productdesc)
		if err != nil {
			panic(err.Error())
		}
		product = append(product, a1)
	}
	json.NewEncoder(w).Encode(product)
}

func main() {

	/*insert, err := db.Query("create table product(productid integer,productname varchar(30),Producttype varchar(20),productdesc varchar(300))")
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
