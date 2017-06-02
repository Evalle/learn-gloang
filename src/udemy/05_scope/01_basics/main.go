package main

import "fmt"

var x int = 42

func main() {
  fmt.Println(x)
  foo()
  y := "some sort of string"
}

func foo()  {
  fmt.Println(x)
  fmt.Println(y)
}
