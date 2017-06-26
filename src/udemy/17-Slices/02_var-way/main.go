package main

import "fmt"

func main() {
	var student []string
	var students [][]string
	// student[0] = "Todd" <- index out of range, because we didn't specify the length of slice
	student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
