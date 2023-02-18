package main

import "fmt"

func main() {
  angka := 5
	for i := angka; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := angka; k >= i; k-- {
			fmt.Print("*")
			fmt.Print(" ")
		}
		fmt.Println()
	}

}
