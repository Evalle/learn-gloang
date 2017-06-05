package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "in the ring"
		fmt.Println(y)
	}
  // you can't access y variable here
}
