package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	WordbyWordScan()
}
func WordbyWordScan() {
	file, err := os.Open("./sample.txt.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
