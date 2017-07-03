package main

import "fmt"

func main() {
	m := map[string]string{
		"Tod":    "teacher",
		"Evgeny": "student",
	}
	fmt.Println(len(m))
	m["Laptop"] = "MacBook"
	fmt.Println(len(m))
}
