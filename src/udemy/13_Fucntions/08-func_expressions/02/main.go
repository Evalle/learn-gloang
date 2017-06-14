package main

import "fmt"

func greeter() func() string {
	return func() string {
		return "Hello world"
	}
}

func main() {
	n := greeter()
	fmt.Println(n())
}
