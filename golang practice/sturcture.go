package main

import "fmt"

type emp struct {
	name    string
	address string
	age     int
}

func display(e emp) {
	fmt.Println(e.name)
}

func main() {
	var empdata1 emp
	empdata1.name = "John"
	empdata1.address = "Street-1, London"
	empdata1.age = 30
	empdata2 := emp{"Raj", "Building-1, Paris", 25}
	display(empdata1)
	display(empdata2)
}
