package main

import "fmt"

func main() {
	m := map[string]string{
		"Tod":    "teacher",
		"Evgeny": "student",
	}
	m["Laptop"] = "Tool"

	fmt.Println(m)
}
