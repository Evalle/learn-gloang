package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
  // don't do that; it's the bad practice to shadow the variables
  max := max(7)
	fmt.Println(max)
}
