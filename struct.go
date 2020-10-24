package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"B", 20})

	fmt.Println(person{name: "A", age: 30})

	fmt.Println(person{name: "F"})

	fmt.Println(&person{name: "S", age: 40})

	fmt.Println(newPerson("J"))

	s := person{name: "R", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
