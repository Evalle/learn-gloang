package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	page, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	fmt.Printf("%s", page)
}
