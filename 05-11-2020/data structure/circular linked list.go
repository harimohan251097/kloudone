package main

import (
	"container/ring"
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := ring.New(len(numbers))

	for i := 0; i < r.Len(); i++ {
		r.Value = numbers[i]
		r = r.Next()
	}

	r.Do(func(x interface{}) {
		fmt.Print(x)
	})
	fmt.Println()

	for i := 0; i < r.Len(); i++ {
		fmt.Print(r.Value)
		r = r.Prev()
	}
	fmt.Println()

	r = r.Move(5)
	r.Do(func(x interface{}) {
		fmt.Print(x)
	})
	fmt.Println()
}
