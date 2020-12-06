// models.article.go

package main

import (
	"database/sql"
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList []article

func addarticle(u article) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		panic(err.Error())
	}
	k := "insert into article5 values( Null,'" + u.Title + "','" + u.Content + "')"
	//{"Articlename":"hari","ID":"3","Atype":"person","Articledesc":"Hari Mohan"}
	result, err := db.Query(k)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
}

func printarticle() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("SELECT * from article5")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var a1 article
		err := result.Scan(&a1.ID, &a1.Title, &a1.Content)
		if err != nil {
			panic(err.Error())
		}
		articleList = append(articleList, a1)
	}
}

// Return a list of all the articles
func getAllArticles() []article {
	articleList = nil
	printarticle()
	return articleList
}

// Fetch an article based on the ID supplied
func getArticleByID(id int) (*article, error) {
	articleList = nil
	printarticle()
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// Create a new article with the title and content provided
func createNewArticle(title, content string) (*article, error) {
	// Set the ID of a new article to one more than the number of articles
	a := article{Title: title, Content: content}
	addarticle(a)
	return &a, nil
}
