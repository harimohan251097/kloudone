package main

import "fmt"

var fac uint64 = 1
var i int = 1
var n int

func fact(n int) uint64 {
	if n < 0 {
		fmt.Print("Factorial  doesn't exist.")
	} else {
		for i := 1; i <= n; i++ {
			fac *= uint64(i)
		}

	}
	return fac
}

func main() {
	fmt.Print("Enter a  integer  ")
	fmt.Scan(&n)
	fmt.Print("Factorial is  ", fact(n))
}
