package main

import "fmt"

func main() {

	c := make(map[string]int)

	c["a"] = 7
	c["b"] = 13

	fmt.Println("map:", c)

	v1 := c["a"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(c))

	delete(c, "b")
	fmt.Println("map:", c)

	_, prs := c["b"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
