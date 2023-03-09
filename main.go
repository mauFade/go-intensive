package main

import (
	"fmt"
	"time"
)

func Workers(workerId int, channel chan int) {
	for x := range channel {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan int)
	workersQuantity := 3

	for i := 0; i < workersQuantity; i++ {
		go Workers(i, channel)
	}

	for i := 0; i < 100000; i++ {
		channel <- i
	}
}
