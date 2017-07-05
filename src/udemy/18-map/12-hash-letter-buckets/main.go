package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split funciton for the scanner operation
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123])
}

// HashBucket get the first letter of the word and convert it into number
func HashBucket(word string) int {
	return int(word[0])
}
