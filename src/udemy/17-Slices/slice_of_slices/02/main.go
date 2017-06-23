package main

import "fmt"

func main() {
	records := make([][]string, 0)
	fmt.Println(records)
	student1 := make([]string,5)
	for i := 0;i < 5;i++ {
		student1[i] = "phil"
	}
	records = append(records, student1)
	fmt.Println(records)
}
