package main

import "fmt"

// HashBucket - return the bucket
func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	fmt.Println(letter)
	bucket := letter % buckets
	return bucket
}

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}
