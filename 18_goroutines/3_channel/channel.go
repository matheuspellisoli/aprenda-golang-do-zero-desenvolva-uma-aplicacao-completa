package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go write("olá mundo!", 5, channel)

	for mes := range channel {
		fmt.Println(mes)
	}
}

func write(text string, count int, channel chan string) {
	for i := 0; i < count; i++ {
		channel <- text
		time.Sleep(time.Second)
	}

	close(channel)
}
