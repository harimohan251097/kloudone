package main

import "fmt"

type point struct {
	x int
	y int
}

type circle struct {
	radius float64
	center *point
}

func main() {

	c1 := circle{4.56, &point{4, 5}}
	fmt.Println("c1.center", c1.center)
	fmt.Println("c1.center.x", c1.center.x)
	fmt.Println("c1.center.y", c1.center.y)
	//fmt.Println("c1.x", c1.x)
	//fmt.Println("c1.y", c1.y)
}
