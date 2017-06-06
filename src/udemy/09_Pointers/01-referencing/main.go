package main

import "fmt"

func main() {
	a := 42
	fmt.Println("a - ", a)
	fmt.Println("mem address of a - ", &a)
	var b *int = &a
	fmt.Println("b - ", b)
	fmt.Println("mem address of b - ", b)
}
