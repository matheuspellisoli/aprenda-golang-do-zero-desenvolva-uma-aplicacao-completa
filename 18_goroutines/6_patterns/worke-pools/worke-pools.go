package main

import "fmt"

func main() {

	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}

	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(tasks <-chan int, results chan<- int) {

	for number := range tasks {
		results <- fiboncci(number)
	}
}

func fiboncci(position int) int {
	if position <= 1 {
		return position
	}

	return fiboncci(position-2) + fiboncci(position-1)
}
