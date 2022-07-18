package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			channel1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"
		}
	}()

	for {
		select {
		case mgChannel1 := <-channel1:
			fmt.Println(mgChannel1)
		case mgChannel2 := <-channel2:
			fmt.Println(mgChannel2)
		}
	}
}
