package main

import (
	"fmt"
)

func main() {
	var messages chan string = make(chan string)
	go func(message string) {
		messages <- message
	}("Goroutines - Ping!")

	fmt.Println(<-messages)
}
