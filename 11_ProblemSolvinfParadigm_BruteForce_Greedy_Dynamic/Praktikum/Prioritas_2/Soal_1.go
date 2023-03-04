package main

import (
	"fmt"
	"math"
)

func Frog(A []int) int {
	n := len(A)
	dp := make([]int, n)

	dp[0] = 0

	for i := 1; i < n; i++ {
		diff1 := dp[i-1] + int(math.Abs(float64(A[i]-A[i-1])))
		var diff2 int
		if i >= 2 {
			diff2 = dp[i-2] + int(math.Abs(float64(A[i]-A[i-2])))
		}
		if diff1 < diff2 {
			dp[i] = diff1
		} else {
			dp[i] = diff2
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         //30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}
