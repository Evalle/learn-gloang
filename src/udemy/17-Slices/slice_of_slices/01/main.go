package main

import "fmt"

func main() {
	records := make([][]string, 0)
	fmt.Println(records)
	student1 := make([]string,5)
	student1[0] = "Phil"
	student1[1] = "Collins"
	student1[2] = "100%"
	student1[3] = "90%"
	student1[4] = "95%"
	records = append(records, student1)
	fmt.Println(records)
	student2 := make([]string,5)
	student2[0] = "Nil"
	student2[1] = "Fry"
	student2[2] = "0%"
	student2[3] = "22%"
	student2[4] = "100%"
	records = append(records, student2)
	fmt.Println(records)
}
