// Write a function, foo, which can be called in all of these ways:
// func main() {
// 	foo(1, 2)
// 	foo(1, 2, 3)
// 	aSlice := []int{1, 2, 3, 4}
// 	foo(aSlice...)
// 	foo()
// }
package main

import "fmt"

func foo(sf ...int) int {
	var biggest int
	for _, v := range sf {
		if v > biggest {
			biggest = v
		}
	}
	fmt.Println(biggest)
	return biggest
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
