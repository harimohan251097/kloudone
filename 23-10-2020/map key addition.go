package main

import (
	"fmt"
)

func main() {
	empsal := map[string]int{
		"hari":  12000,
		"mohan": 15000,
	}
	newEm := "hari"
	val, ok := empsal[newEm]
	if ok == true {
		fmt.Println("Salary of", newEm, "is", val)
		return
	}
	fmt.Println(newEm, "not found")
}
