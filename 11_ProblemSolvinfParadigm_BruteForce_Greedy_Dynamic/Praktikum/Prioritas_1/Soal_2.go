package main

import (
	"fmt"
)

func generatePascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
		triangle[i] = row
	}
	return triangle
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(generatePascalTriangle(n))

}
