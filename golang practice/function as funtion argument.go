package main

import "fmt"

func main() {

	x := 3
	y := 4

	add := getAddSub()
	r1 := apply(x, y, add)

	fmt.Printf("%d + %d = %d\n", x, y, r1)

}

func apply(x, y int, add func(int, int) int) int {

	r1 := add(x, y)

	return r1
}

func getAddSub() func(int, int) int {

	add := func(x, y int) int {
		return x + y
	}

	return add
}
