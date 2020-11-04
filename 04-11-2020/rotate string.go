package main

import "fmt"

type Runes []rune

func (str Runes) ReverseString() (revStr Runes) {
	l := len(str)
	revStr = make(Runes, l)
	for i := 0; i <= l/2; i++ {
		revStr[i], revStr[l-1-i] = str[l-1-i], str[i]
	}
	return revStr
}

func (str Runes) String() string {
	return string(str)
}

func main() {
	inputStr := "Most Popular Golang String Functions"
	strRune := Runes(inputStr)
	outputStr := strRune.ReverseString()
	fmt.Println("Original: ", inputStr)
	fmt.Println("Reversed: ", outputStr)
}
