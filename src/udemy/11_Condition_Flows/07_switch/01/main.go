package main

import "fmt"

func main() {
	switch "Mei" {
	case "Tim":
		fmt.Println("Hello Tim")
	case "Jim":
		fmt.Println("Hello Jim")
	case "Medhi":
		fmt.Println("Hello Medhi")
	default:
		fmt.Println("You have nil friends?")
	}
}
