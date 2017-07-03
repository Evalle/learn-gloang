package main

import "fmt"

func main() {
	m := map[string]string{
		"Tod":    "teacher",
		"Evgeny": "student",
	}
	fmt.Println(m)
	delete(m, "Evgeny")
	fmt.Println(m)
}
