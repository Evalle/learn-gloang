package main

import "fmt"

func main() {
  for i := 0;i < 10;i++ {
    if i % 2 == 0 {
      if i != 0 {
      fmt.Println(i, "even")
    }
    }
  }
}
