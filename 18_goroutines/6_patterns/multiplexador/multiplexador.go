package main

import (
	"fmt"
	"time"
)

func main() {

	channel := multiplexador(write("olá mundo"), write("teste"))

	for msg := range channel {
		fmt.Println(msg)
	}

}

func multiplexador(channelInput1, channelInput2 <-chan string) <-chan string {

	channelOutput := make(chan string)

	go func() {
		for {
			select {
			case msg := <-channelInput1:
				channelOutput <- msg
			case msg := <-channelInput2:
				channelOutput <- msg
			}
		}
	}()

	return channelOutput

}

func write(text string) <-chan string {

	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor [%s]", text)
			time.Sleep(time.Second)
		}
	}()

	return channel
}
