package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeX(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	i := 2
	for {
		if isPrime(i) {
			count++
			if count == n {
				return i
			}
		}
		i++
	}
}

func main() {

	fmt.Println(primeX(1))  //2
	fmt.Println(primeX(5))  //11
	fmt.Println(primeX(8))  //19
	fmt.Println(primeX(9))  //23
	fmt.Println(primeX(10)) //29

}
