package main

import "fmt"

func main() {

	x := 7
	y := &x
	fmt.Println("Before: ", x, y)
	*y = 8
	fmt.Println("Before: ", x, y)
}
