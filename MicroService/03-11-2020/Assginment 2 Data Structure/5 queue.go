package main

import "fmt"

func push(que []int, element int) []int {
	que = append(que, element)
	fmt.Println("push:", element)
	return que
}

func pop(que []int) []int {
	element := que[0]
	fmt.Println("pop:", element)
	return que[1:]
}

func main() {
	var que []int

	que = push(que, 10)
	que = push(que, 20)
	que = push(que, 30)

	fmt.Println("Queue:", que)

	que = pop(que)
	fmt.Println("Queue:", que)

	que = push(que, 40)
	fmt.Println("Queue:", que)
}
