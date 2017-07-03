package main

import "fmt"

func main() {
	m := map[int]string{
		1: "Hello",
		2: "Bonjour",
		3: "Ahoj",
	}
	fmt.Println(m)
	delete(m, 2)
	fmt.Println(m)
}
