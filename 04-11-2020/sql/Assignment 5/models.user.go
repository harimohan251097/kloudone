// models.user.go

package main

import (
	"database/sql"
	"errors"
	"strings"

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
