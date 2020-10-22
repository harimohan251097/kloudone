package main

import (
	"fmt"
	"strconv"
)

func FiboRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FiboRecursion(n-1) + FiboRecursion(n-2)
}

func main() {
	var k int
	fmt.Print("Enter a  integer  ")
	fmt.Scan(&k)
	fmt.Println("")
	for i := 0; i <= k; i++ {
		fmt.Print(strconv.Itoa(FiboRecursion(i)) + " ")
	}
	fmt.Println("")
}
