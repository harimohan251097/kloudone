package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Product struct {
	Productid   string
	Productname string
	Producttype string
	Productdesc string
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
		err := result.Scan(&a1.Productid, &a1.Productname, &a1.Producttype, &a1.Productdesc)
		if err != nil {
			panic(err.Error())
		}

		product = append(product, a1)
	}
	fmt.Print(product)
	json.NewEncoder(w).Encode(product)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/product", GetProductEndpoint).Methods("GET")
	http.ListenAndServe(":123", router)
}
