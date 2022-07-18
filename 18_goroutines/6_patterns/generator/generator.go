package main

import (
	"fmt"
	"time"
)

func main() {

	channel := write("Teste")

	for msg := range channel {
		fmt.Println(msg)
	}

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
