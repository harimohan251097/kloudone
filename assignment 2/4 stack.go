package main

import (
	"fmt"
)

type item struct {
	value interface{}
	next  *item
}

type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() (value interface{}) {
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return
	}

	return nil
}

func main() {
	stack := new(Stack)
	// Push different data type to the stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push("A")
	stack.Push("B")
	stack.Push("C")
	stack.Push("D")

	var head = stack.Len()
	//fmt.Println(head)
	for stack.Len() > 0 {
		k := stack.Pop()
		if stack.Len() == head-1 {
			fmt.Printf("head is ")
			fmt.Println(k)
			fmt.Println("\n\n Elemnt of stack is : ")
		}
		fmt.Println(k)
		if stack.Len() == 0 {
			fmt.Println(" ")
			fmt.Printf("tail is ")
			fmt.Println(k)
		}
	}
}
