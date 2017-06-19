// Write a function with one variadic parameter that finds the greatest number
// in a list of numbers.
//
// [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.
package main

import "fmt"

func findBiggest(sf ...int) int {
	var biggest int
	for _, v := range sf {
		fmt.Println(v)
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}

func main() {
	n := findBiggest(2, 4, 6, 8)
	fmt.Println(n)
}
