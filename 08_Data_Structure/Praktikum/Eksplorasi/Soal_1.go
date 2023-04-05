package main

import (
	"fmt"
	"math"
)

func AbosoluteMatriks(matriks [][]int) int {
	var diagonalKiri, diagonalKanan int
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {
			if i == j {
				diagonalKiri += matriks[i][j]
			}
			if i+j == len(matriks)-1 {
				diagonalKanan += matriks[i][j]
			}
		}
	}
	fmt.Println("jumlah diagonal kiri: ", diagonalKiri)
	fmt.Println("jumlah diagonal kanan: ", diagonalKanan)

	//nilai absolut
	var nilaiAbsolut int
	nilaiAbsolut = diagonalKiri - diagonalKanan
	return int(math.Abs(float64(nilaiAbsolut)))
}

func main() {
	fmt.Println(AbosoluteMatriks([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}))

}
