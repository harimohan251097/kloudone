package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Welcome to Geeks for Geeks"
	str2 := "Here! we learn about go strings"

	fmt.Println("Original strings")
	fmt.Println("String 1: ", str1)
	fmt.Println("String 2: ", str2)

	res1 := strings.Contains(str1, "Geeks")
	res2 := strings.Contains(str2, "GFG")

	fmt.Println("\nResult 1: ", res1)
	fmt.Println("Result 2: ", res2)

}
