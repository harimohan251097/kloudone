package main

import "fmt"

func main() {

	a := [5]string{"one", "two", "three", "four", "five"}
	fmt.Println("Array after creation:", a)

	var b []string = a[1:4]
	fmt.Println("Slice after creation:", b)

	b[0] = "changed"
	fmt.Println("Slice after modifying:", b)
	fmt.Println("Array after slice modification:", a)

}
