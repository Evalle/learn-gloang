package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"one":   "Hello",
		"two":   "Hi",
		"three": "Ahoj",
	}
	fmt.Println(myGreeting)
	for key, value := range myGreeting {
		fmt.Println(key, "-", value)
	}
}
