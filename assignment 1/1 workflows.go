package main

import (
	"fmt"
	"sync"
)

type Course struct {
	Name string
}
type employee struct {
	Name string
}
type RegisterStudentsResults struct {
	Results []StudentRegistrationResult
}
type StudentRegistration struct {
	Student employee
	Course  Course
}
type StudentRegistrationResult struct {
	Registration StudentRegistration
	Error        error
}

func RegisterStudents(students []employee, course Course) RegisterStudentsResults {
	output := make(chan RegisterStudentsResults)
	input := make(chan StudentRegistrationResult)
	var wg sync.WaitGroup
	go handleResults(input, output, &wg)
	defer close(output)
	for _, student := range students {
		wg.Add(1)
		go ConcurrentRegisterStudent(student, course, input)
	}

	wg.Wait()
	close(input)
	return <-output
}

func handleResults(input chan StudentRegistrationResult, output chan RegisterStudentsResults, wg *sync.WaitGroup) {
	var results RegisterStudentsResults
	for result := range input {
		results.Results = append(results.Results, result)
		wg.Done()
	}
	output <- results
}

func ConcurrentRegisterStudent(student employee, course Course, output chan StudentRegistrationResult) {
	result := RegisterStudent(student, course)
	output <- result
}

func RegisterStudent(student employee, course Course) StudentRegistrationResult {
	return StudentRegistrationResult{
		Registration: StudentRegistration{
			Student: student,
			Course:  course,
		},
	}
}
func main() {
	students := []employee{
		employee{Name: "mam"},
		employee{Name: "Hari"},
	}
	result := RegisterStudents(students, Course{Name: "program"})
	fmt.Print(result)
}
