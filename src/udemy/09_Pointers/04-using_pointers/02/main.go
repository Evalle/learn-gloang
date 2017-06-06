package main

import "fmt"

func zero(z int) {
	fmt.Println("mem address in the zero function", &z) // mem address in the zero fucntion
	z = 0
}

func main() {
	x := 5
	fmt.Println("mem address in the main function", &x) // mem address in the main function
	zero(x)
	fmt.Println(x)
  fmt.Println(&x)
}
