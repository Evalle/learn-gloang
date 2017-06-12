package main

import "fmt"

func main() {
	var (
    smallNumber int
    largeNumber int
  )
	fmt.Println("Enter small number")
	fmt.Scan(&smallNumber)
	fmt.Println("Enter large number")
	fmt.Scan(&largeNumber)
	fmt.Println("The Remainder is", largeNumber % smallNumber)
}
