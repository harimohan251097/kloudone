package main

import (
	"bufio"
	"fmt"
)

func main() {
	fmt.Println("enter no.(10-100)")
	Scanner := bufio.Newscanner(OS.stdin)
	Scanner.scan()
	x := Scanner.Text()
	if x < 10 {
		fmt.Println("x is less than 10")
	} else if x >= 10 && x <= 90 {

		fmt.Println("x is between 10 and 90")
	} else {
		fmt.Println("x is greater than 90")
	}
}
