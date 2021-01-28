// models.user.go

package main

import (
	"math/rand"
	_ "math/rand"
	"net/http"
	_ "net/http"
	"strconv"
	_ "strconv"

	"database/sql"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList []user

func adduser(u user) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		panic(err.Error())
	}
	k := "insert into user5 values('" + u.Username + "','" + u.Password + "')"
	//{"Articlename":"hari","ID":"3","Atype":"person","Articledesc":"Hari Mohan"}
	result, err := db.Query(k)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
}

func printuser() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query("SELECT * from user5")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var a1 user
		err := result.Scan(&a1.Username, &a1.Password)
		if err != nil {
			panic(err.Error())
		}
		userList = append(userList, a1)
	}
}

// Check if the username and password combination is valid
func isUserValid(username, password string) bool {
	printuser()
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func registerNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := user{Username: username, Password: password}
	adduser(u)

	return &u, nil
}

// Check if the supplied username is available
func isUsernameAvailable(username string) bool {
	printuser()
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}

func showLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func performLogin(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	//var sameSiteCookie http.SameSite

	// Check if the username/password combination is valid
	if isUserValid(username, password) {
		// If the username/password is valid set the token in a cookie
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {

		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func generateSessionToken() string {

	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {

	//var sameSiteCookie http.SameSite

	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func showRegistrationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	//var sameSiteCookie http.SameSite

	if _, err := registerNewUser(username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful registration & Login"}, "login-successful.html")

	} else {

		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})

	}
}
