package main

import "fmt"

func main() {
	i, j := 2, 1
	switch i + j {
	case 1:
		fmt.Println("Sum is 1")
	case 2:
		fmt.Println("Sum is 2")
	case 3:
		fmt.Println("Sum is 3")
	default:
		fmt.Println("Printing default")
	}
}
