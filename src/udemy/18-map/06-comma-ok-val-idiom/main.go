package main

import "fmt"

func main() {
	m := map[int]string{
		1: "Hi",
		2: "Hello",
		3: "Bobjour",
		4: "Ahoj",
	}
	fmt.Println(m)
	if val, exists := m[2]; exists {
		fmt.Println("That value exists")
		fmt.Println("val", val)
		fmt.Println("exists", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val", val)
		fmt.Println("exists", exists)
	}
}
