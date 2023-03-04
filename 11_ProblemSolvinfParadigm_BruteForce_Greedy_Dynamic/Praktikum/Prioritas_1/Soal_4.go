package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	// x + y + z = a
	// x * y * z = b
	// x ^ 2 + y ^ 2 + z ^ 2 = c

	// mencari nilai x
	for x := 1; x*x <= 10000; x++ {
		for y := 0; y*y <= 10000; y++ {
			for z := 0; z*z <= 10000; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Println(x, y, z)
				}
			}
		}
	}

	fmt.Println("no solution")

}

func main() {

	SimpleEquations(1, 2, 1)
	SimpleEquations(6, 6, 14)

}
