package main

import (
	"fmt"
)

func getBinaryArray(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		for j := i; j > 0; j >>= 1 {
			if j&1 == 1 {
				count++
			}
		}
		ans[i] = count
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(getBinaryArray(n))
}
