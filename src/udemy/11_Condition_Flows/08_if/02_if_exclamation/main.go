package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This did no run")
	}
	if !false {
		fmt.Println("This ran")
	}
}
