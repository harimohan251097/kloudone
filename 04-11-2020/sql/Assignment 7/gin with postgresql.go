package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Product struct {
	Productid   string `json:"Productid,omitempty"`
	Productname string `json:"Productname,omitempty"`
	Producttype string `json:"Producttype,omitempty"`
	Productdesc string `json:"Productdesc,omitempty"`
}

func main() {
	r := gin.Default()
	r.GET("/product", GetProductEndpoint)
	r.GET("/product/:id", GetProduct1Endpoint)
	r.POST("/product", CreateProductEndpoint)
	r.PUT("/product/:id", PutProductEndpoint)
	r.DELETE("/product/:id", DeleteProductEndpoint)
	r.Run(":123")
}

func DeleteProductEndpoint(c *gin.Context) {
	id := c.Params.ByName("id")
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	query := "delete from product where Productid=" + id
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

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
	c.JSON(200, product)
}

func PutProductEndpoint(c *gin.Context) {
	id := c.Params.ByName("id")
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	k := "update product set Productname='" + keyVal["Productname"] + "',Producttype='" + keyVal["Producttype"] + "',Productdesc='" + keyVal["Productdesc"] + "' where Productid=" + id
	//{"Productname":"hari123","Producttype":"person123","Productdesc":"Hari Mohan"}
	result, err := db.Query(k)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
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
	c.JSON(200, product)
}

func CreateProductEndpoint(c *gin.Context) {
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(c.Request.Body)
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
	c.JSON(200, product)
}

func GetProduct1Endpoint(c *gin.Context) {
	id := c.Params.ByName("id")
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	query := "SELECT * from product WHERE Productid =" + id
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
	c.JSON(200, a1)
}

func GetProductEndpoint(c *gin.Context) {
	connStr := "user=postgres dbname=hari password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)

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
	c.JSON(200, product)
}
