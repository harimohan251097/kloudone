package main

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FName, LName, Email string
	Age                 int
	MonthlySalary       []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FName, e.LName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("===================")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "----------------------"
}

func main() {

	e := Employee{FName: "A", LName: "B", Email: "AB@gmail.com", Age: 2,
		MonthlySalary: []Salary{Salary{Basic: 15000.00, HRA: 5000.00, TA: 2000.00},
			Salary{Basic: 16000.00, HRA: 5000.00, TA: 2100.00},
			Salary{Basic: 17000.00, HRA: 5000.00, TA: 2200.00},
		},
	}
	fmt.Println(e.EmpInfo())
}
