package main

import "fmt"

func main() {
	var x []int = []int{4, 5, 6}
	y := x
	y[0] = 100
	fmt.Println("Before: ", x)
	fmt.Println("After: ", y)
}
