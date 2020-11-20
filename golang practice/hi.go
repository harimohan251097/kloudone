package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func postHomePage(c *gin.Context) {
	body := c.Request.Body // to read the body
	value, err := ioutil.ReadAll(body)

	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message": string(value), // entering the message
	})
}

func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello KloudOne!",
	})
}

// function to parse values
func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// same as QueryStrings function, it just uses param instead
func PathParameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("Hello KloudOne!")

	r := gin.Default()
	r.GET("/", homePage)
	r.POST("/", postHomePage)
	r.GET("/query", QueryStrings)             // /query?name=mansi&age=19
	r.GET("/path/:name/:age", PathParameters) // /path/mansi/18/
	r.Run()
}
