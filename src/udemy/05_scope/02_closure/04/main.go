package main

import "fmt"

var x int = 0

func increment() int {
	x++
	return x
}

func main() {
	increment()
	increment()
	fmt.Println(x)
}
