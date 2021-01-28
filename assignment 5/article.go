// models.article.go

package main

import (
	"net/http"
	_ "net/http"
	"strconv"
	_ "strconv"

	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
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

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func showArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

func getArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := getArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func createArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(title, content); err == nil {
		// If the article is created successfully, show success message
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
