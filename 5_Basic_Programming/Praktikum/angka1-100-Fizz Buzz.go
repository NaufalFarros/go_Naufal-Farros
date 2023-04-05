package main

import (
	"fmt"
)

func main() {
for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
			fmt.Print(" ")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
			fmt.Print(" ")
		} else {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}

}
