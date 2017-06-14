package main

import "fmt"

func greeting(name string) {
	fmt.Println("Hello", name)
}

func main() {
	greeting("Marta")
	greeting("Stewart")
}

// greeting is declared with a parametr
// when calling greeting, pass in an argument
