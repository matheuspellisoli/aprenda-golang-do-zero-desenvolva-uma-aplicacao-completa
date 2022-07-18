package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "ola"
	channel <- "ol2a"

	message := <-channel
	message2 := <-channel
	fmt.Println(message, message2)

}
