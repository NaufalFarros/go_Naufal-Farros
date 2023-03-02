package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (int, int) {
	//mencari nilai min dan max menggunakan pointer dan print pointer
	var min, max int = *numbers[0], *numbers[0]

	//mengetahui pointer
	// fmt.Println("pointer numbers :", numbers)
	// fmt.Println("pointer min :", &min)
	// fmt.Println("pointer max :", &max)

	for _, number := range numbers {
		// fmt.Println("pointer number :", number)
		if *number < min {
			min = *number
		}
		if *number > max {
			max = *number
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai Min :", min)
	fmt.Println("Nilai Max :", max)
}
