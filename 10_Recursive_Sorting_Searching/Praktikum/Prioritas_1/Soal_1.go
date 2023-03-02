package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		// cetak perhitungan fibonacci
		// fmt.Println("fibonacci(", n, ") = fibonacci(", n-1, ") + fibonacci(", n-2, ")")
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144

}
