package main

import (
	"fmt"
	"sync"
)

func calculateFactorial(num int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	fmt.Printf("Faktorial dari %d adalah %d\n", num, result)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	nums := []int{5, 6, 7, 8, 9}

	for _, num := range nums {
		wg.Add(1)
		go calculateFactorial(num, &wg, &mu)
	}

	wg.Wait()
}
