package main

import "fmt"

func main() {
	mySlice := []int{1,3,5,7}
	myOtherslice := []int{9,11,13,15}
	mySlice = append(mySlice, myOtherslice...)
	fmt.Println(mySlice)
}
