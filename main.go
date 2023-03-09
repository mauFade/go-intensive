package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ": ", i)

		time.Sleep(time.Second)
	}
}

func main() {
	go task("Task 1") // O comando "go" inicia uma nova thread
	go task("Task 2")
	task("Task 3")

	channel := make(chan string) // Canal de comunicação entre as threads
	numChannel := make(chan int)

	// Thread 2
	go func() {
		// Inicia o programa, cria um canal de cominuacação entre thrads e passa um valor de uma thread pra outra
		channel <- "Hello new thread 2\nDentro da thread 2 passei um dado pra thread 1\n"
	}()

	go Publish(numChannel)

	go Reader(numChannel)

	time.Sleep(time.Second * 2)
}

func Reader(channel chan int) {
	for value := range channel {

		fmt.Println(value)
	}
}

func Publish(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}

	close(channel)
}
