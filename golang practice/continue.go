package main

import "fmt"

func main() {
	for k := 0; k < 10; k++ {
		if k == 5 {
			fmt.Println("Continue loop")
			continue
		}
		fmt.Println("k is", k)
	}
	fmt.Println("Exiting program")
}
