package main

import (
	"fmt"
)

func printMultiples(ch chan int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			ch <- i
		}
	}
}

func main() {
	ch := make(chan int, 10)
	go printMultiples(ch)

	fmt.Println("Bilangan kelipatan buffer channel 3:")
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}
