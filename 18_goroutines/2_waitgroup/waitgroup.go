package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		write("olá mundo!", 5)
		waitGroup.Done()
	}()

	go func() {
		write("Programando em GO!", 10)
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func write(text string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
