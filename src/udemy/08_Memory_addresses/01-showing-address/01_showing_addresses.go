package main

import "fmt"

var a int = 42

func main() {
	fmt.Println("a - ", a)
	fmt.Println("memory address of a - ", &a)
}
