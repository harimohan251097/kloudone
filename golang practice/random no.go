package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go RandNum(ch)
	num := <-ch
	fmt.Print(num)
}

func RandNum(ch chan int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	g1 := rand.New(s1)
	result := g1.Intn(1000)
	ch <- result
}
