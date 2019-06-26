/**
*   Program making use of goroutine and channel.
*   This is the simple implementation of goroutines with channels.
*   GoDoc Reference: https://golang.org/doc/effective_go.html#goroutines.
**/
package main

import "fmt"

func hello(c chan string) {
	//Receiving value from the channel
	val := <-c
	fmt.Println("Hello, ", val)
}

func main() {
	messages := make(chan string)

	// Start a goroutine
	go hello(messages)

	// Sending value to the channel
	messages <- "world"

	// Used to stop the conteol at the console so
	// we can see the output.
	fmt.Scanln()
}
