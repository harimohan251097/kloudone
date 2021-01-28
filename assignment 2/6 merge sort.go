package main

import (
	"fmt"
)

func MergeSort2(numbers []int, resultChan chan []int) {
	l := len(numbers)
	if l <= 1 {
		resultChan <- numbers
		return
	}

	m := l / 2

	lchan := make(chan []int, 1)
	rchan := make(chan []int, 1)

	go MergeSort2(numbers[0:m], lchan)
	go MergeSort2(numbers[m:l], rchan)
	go MergeSort3(<-lchan, <-rchan, resultChan)
}

func MergeSort3(left []int, right []int, resultChannel chan []int) {
	leftLength := len(left)
	rightLength := len(right)

	if leftLength == 0 {
		resultChannel <- right
		return
	}
	if rightLength == 0 {
		resultChannel <- left
		return
	}

	result := make([]int, (leftLength + rightLength))
	li := 0
	ri := 0
	resulti := 0
	var rnum, lnum int

	for li < leftLength || ri < rightLength {
		if li < leftLength && ri < rightLength {
			lnum = left[li]
			rnum = right[ri]

			if lnum <= rnum {
				result[resulti] = lnum
				li++
			} else {
				result[resulti] = rnum
				ri++
			}

		} else if li < leftLength {
			lnum = left[li]
			result[resulti] = lnum
			li++
		} else if ri < rightLength {
			rnum = right[ri]
			result[resulti] = rnum
			ri++
		}

		resulti++
	}

	resultChannel <- result
}

func Mergesort1() {
	var largeArr []int = []int{5, 3, 4, 1, 2}

	fmt.Println("Unsorted List")
	fmt.Println(largeArr)

	fmt.Println("Merge sort with Goroutines")
	resultChan := make(chan []int, 1)
	MergeSort2(largeArr, resultChan)
	k := <-resultChan
	fmt.Println(k)
}

func main() {
	Mergesort1()
}
