package main

import "fmt"

func main() {
	a := 42
	fmt.Println("a - ", a)                   // 43
	fmt.Println("mem address of a - ", &a)   // some mem address
	var b *int = &a                          // referencing
	fmt.Println("mem address of b - ", b)    // some mem address
	fmt.Println("b - ", *b)                  // dereferencing, we will have the actual value
}
