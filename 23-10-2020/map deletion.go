package main

import "fmt"

func main() {
	var Exten = map[string]string{
		"Py":   ".py",
		"C":    ".cpp",
		"Java": ".java",
		"Go":   ".go",
		"Kot":  ".kt",
	}

	fmt.Println(Exten)

	delete(Exten, "Kot")

	// delete function doesn't do anything if the key doesn't exist
	delete(Exten, "Javascript")

	fmt.Println(Exten)
}
