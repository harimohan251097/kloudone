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

type Article struct {
	ID          string `json:"id,omitempty"`
	Articlename string `json:"Articlename,omitempty"`
	Atype       string `json:"Atype,omitempty"`
	Articledesc string `json:"Articledesc,omitempty"`
}

func GetArticleEndpoint(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	var articles []Article
	result, err := db.Query("SELECT * from article")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var a1 Article
		err := result.Scan(&a1.ID, &a1.Articlename, &a1.Atype, &a1.Articledesc)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, a1)
	}
	fmt.Print(articles)
	json.NewEncoder(w).Encode(articles)
}

func GetArticle1Endpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("SELECT * from article WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var a1 Article
	for result.Next() {
		err := result.Scan(&a1.ID, &a1.Articlename, &a1.Atype, &a1.Articledesc)
		if err != nil {
			panic(err.Error())
		}

	}
	json.NewEncoder(w).Encode(a1)
}

func CreateArticleEndpoint(w http.ResponseWriter, req *http.Request) {
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
	k := "insert into article values(" + keyVal["ID"] + ",'" + keyVal["Articlename"] + "','" + keyVal["Atype"] + "','" + keyVal["Articledesc"] + "')"
	//{"Articlename":"hari","ID":"3","Atype":"person","Articledesc":"Hari Mohan"}
	result, err := db.Query(k)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	fmt.Fprintf(w, "New post was created")

	var articles []Article
	result1, err := db.Query("SELECT * from article")
	if err != nil {
		panic(err.Error())
	}
	defer result1.Close()

	for result1.Next() {
		var a1 Article
		err := result1.Scan(&a1.ID, &a1.Articlename, &a1.Atype, &a1.Articledesc)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, a1)
	}
	json.NewEncoder(w).Encode(articles)
}

func DeleteArticleEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("delete from article where ID=?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	fmt.Fprintf(w, "post is deleted")

	var articles []Article
	result1, err := db.Query("SELECT * from article")
	if err != nil {
		panic(err.Error())
	}
	defer result1.Close()

	for result1.Next() {
		var a1 Article
		err := result1.Scan(&a1.ID, &a1.Articlename, &a1.Atype, &a1.Articledesc)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, a1)
	}
	json.NewEncoder(w).Encode(articles)
}

func PutArticleEndpoint(w http.ResponseWriter, req *http.Request) {
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
	k := "update article set Articlename='" + keyVal["Articlename"] + "',type='" + keyVal["Atype"] + "',Articledesc='" + keyVal["Articledesc"] + "' where ID=" + params["id"]
	//{"Articlename":"hari123","Atype":"person123","Articledesc":"Hari Mohan"}
	result, err := db.Query(k)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	fmt.Fprintf(w, " post was updated")

	var articles []Article
	result1, err := db.Query("SELECT * from article")
	if err != nil {
		panic(err.Error())
	}
	defer result1.Close()

	for result1.Next() {
		var a1 Article
		err := result1.Scan(&a1.ID, &a1.Articlename, &a1.Atype, &a1.Articledesc)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, a1)
	}
	json.NewEncoder(w).Encode(articles)
}

func main() {

	/*insert, err := db.Query(" create table article(ID integer,Articlename varchar(20),type varchar(30),Articledesc varchar(100))")
	insert, err := db.Query(" insert into article values(0,'machine','tech','machine learning')")
	insert, err := db.Query(" insert into article values(1,'welding','machnical','gas welding')")
	*/
	router := mux.NewRouter()
	router.HandleFunc("/article", GetArticleEndpoint).Methods("GET")
	router.HandleFunc("/article/{id}", GetArticle1Endpoint).Methods("GET")
	router.HandleFunc("/article/{id}", PutArticleEndpoint).Methods("PUT")
	router.HandleFunc("/article", CreateArticleEndpoint).Methods("POST")
	router.HandleFunc("/article/{id}", DeleteArticleEndpoint).Methods("DELETE")
	http.ListenAndServe(":123", router)
}
