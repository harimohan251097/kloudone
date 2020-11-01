package main

import "fmt"

func main() {
	var sample []int = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	fmt.Println("Before Sorting")
	for _, val := range sample {
		fmt.Println(val)
	}
	insertionSort(sample)
}

func insertionSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	fmt.Println("After Sorting")
	for _, val := range arr {
		fmt.Println(val)
	}
}
