package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println(isValidUrl("http://www.golangcode.com"))

	fmt.Println(isValidUrl("golangcode.com"))

	fmt.Println(isValidUrl(""))
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
