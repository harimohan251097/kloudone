package main

import "fmt"

func main() {
	a := 20

	var b *int = &a
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of a:", a)
	fmt.Println("Address of pointer b:", b)
	fmt.Println("Value of pointer b", *b)
	fmt.Println("Value of pointer b", *b)
	fmt.Println("Value of a:", a)
}
