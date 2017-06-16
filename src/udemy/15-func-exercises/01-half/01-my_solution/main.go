// Write a function which takes an integer. The function will have two returns.
// The first return should be the argument divided by 2. The second return
// should be a bool that let’s us know whether or not the argument was even.
// For example:
// - half(1) returns (0, false)
// - half(2) returns (1, true)

package main

import "fmt"

func Half(arg int) (int, bool) {
	div := arg / 2
	var sndReturn bool
	if div % 2 == 0 {
		sndReturn = false
	} else {
		sndReturn = true
	}
	return div, sndReturn
}

func main() {
	fmt.Println(Half(1))
}
