package main

import "fmt"

func zero(z *int) {
  fmt.Println("zero function", z)
  *z = 0
}

func main() {
	x := 5
  fmt.Println("main", &x)
	zero(&x)
	fmt.Println("main", x)
}
