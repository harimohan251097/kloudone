package main

import "fmt"

func main() {
	var a int = 1

	for true {
		var n int
		fmt.Printf("value of a: %d\n", a)
		a++
		fmt.Print("Enter \n 1 for continue print \n 0 for stop print \n ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
	}
}
