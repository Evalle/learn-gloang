package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	hourOftheDay := time.Now().Hour()
	greeting, err := getGreeting(hourOftheDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string
	if hour < 7 {
		err := errors.New("Too early for greetings")
		return message, err
	}
	if hour < 12 {
		message = "Good morning!"
	} else if hour < 18 {
		message = "Good afternoon!"
	} else {
		message = "Good evening!"
	}
	return message, nil
}
