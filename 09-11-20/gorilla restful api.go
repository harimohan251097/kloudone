package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	Articles := []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	//myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)

	log.Fatal(http.ListenAndServe(":12345", myRouter))
}

func main() {

	handleRequests()
}
