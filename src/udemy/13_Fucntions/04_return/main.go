package main

import "fmt"

func greeting(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func main() {
	fmt.Println(greeting("Jane", "Doe"))
}
