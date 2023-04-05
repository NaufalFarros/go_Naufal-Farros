package main

import (
	"fmt"
	"sync"
)

func countAlphabet(text string, result chan map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()

	letterCount := make(map[rune]int)
	for _, c := range text {
		letterCount[c]++
	}

	result <- letterCount
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	result := make(chan map[rune]int)
	var wg sync.WaitGroup

	n := len(text) / 4
	wg.Add(4)

	go countAlphabet(text[:n], result, &wg)
	go countAlphabet(text[n:2*n], result, &wg)
	go countAlphabet(text[2*n:3*n], result, &wg)
	go countAlphabet(text[3*n:], result, &wg)

	go func() {
		wg.Wait()
		close(result)
	}()

	mergedResult := make(map[rune]int)
	for letterCount := range result {
		for letter, count := range letterCount {
			mergedResult[letter] += count
		}
	}

	for c, count := range mergedResult {
		fmt.Printf("%c: %d\n", c, count)
	}
}
