package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(mySlice)
	fmt.Println("[1:2] - ", mySlice[1:3])
	fmt.Println("[4:] - ", mySlice[4:])
}
