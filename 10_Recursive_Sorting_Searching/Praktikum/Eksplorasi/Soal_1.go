package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	maxAwal := arr[0]
	maxTerakhir := arr[0]
	// fmt.Println(maxAwal, maxTerakhir)
	for i := 1; i < len(arr); i++ {
		maxTerakhir = max(arr[i], maxTerakhir+arr[i])
		maxAwal = max(maxAwal, maxTerakhir)
	}

	return maxAwal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
}
