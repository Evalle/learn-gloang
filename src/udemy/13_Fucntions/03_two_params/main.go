package main

import "fmt"

func greeting(name, sndname string) {
	fmt.Println(name, sndname)
}

func main() {
	greeting("Alice", "Bob")
}
