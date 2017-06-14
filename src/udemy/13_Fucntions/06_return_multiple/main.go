package main

import "fmt"

func greeting(name, sndname string) (string, string) {
	return fmt.Sprint(name, sndname), fmt.Sprint(sndname, name)
}

func main() {
	fmt.Println(greeting("Alice ", "Bob "))
}
