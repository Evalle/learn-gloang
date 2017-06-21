package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "d", "e", "f", "g"} // slice
	fmt.Println(mySlice)
	fmt.Println(mySlice[0])
	fmt.Println(mySlice[0:2])
	fmt.Println(mySlice[0:])
	fmt.Println("mySlice"[2])
}
