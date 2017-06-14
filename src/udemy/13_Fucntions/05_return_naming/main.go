package main

import "fmt"

func greeting(name, sndname string) (s string) {
	s = fmt.Sprint(name, sndname)
	return
}

func main() {
	fmt.Println(greeting("Alice", "Bob"))
}
