package main

import "fmt"

func main() {
	student := make([]string, 35)
	// We specify the length of the slice (35)
	students := make([][]string, 35)
	// so we can add a new student without appending
	student[0] = "Todd"
	fmt.Println(student)
	students = append(students, student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
