package main

import "fmt"

func main() {
	mySlice := []int{1,3,5,7,9}
	var x [10]int
	fmt.Println(mySlice, x)
	fmt.Printf("%T, %T", mySlice, x)
}
