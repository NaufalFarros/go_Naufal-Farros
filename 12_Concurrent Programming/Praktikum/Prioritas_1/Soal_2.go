package main

import (
	"fmt"
)

func printMultiples(channel chan int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			channel <- i
		}
	}
}

func main() {
	channel := make(chan int)
	// fmt.Println("Channel created", ch)
	go printMultiples(channel)

	fmt.Println("Bilangan kelipatan 3:")
	for i := 1; i <= 10; i++ {
		fmt.Println(<-channel)
	}
}
