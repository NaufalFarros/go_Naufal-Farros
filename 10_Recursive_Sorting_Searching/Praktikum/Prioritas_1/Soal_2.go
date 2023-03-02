package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(arr []string) []pair {
	// buat map untuk menyimpan hasil
	result := make(map[string]int)

	for i := 0; i < len(arr); i++ {
		if _, ok := result[arr[i]]; ok {
			result[arr[i]] += 1

		} else {
			result[arr[i]] = 1
		}
	}

	bublesort := make([]pair, len(result))
	i := 0
	for k, v := range result {
		bublesort[i] = pair{k, v}
		i++
	}

	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if bublesort[i].count > bublesort[j].count {
				temp := bublesort[i]
				bublesort[i] = bublesort[j]
				bublesort[j] = temp
			}
		}
	}
	return bublesort
}

func main() {

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	// golang -> 1, ruby -> 2, js -> 4
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	// C->1, D->2, B->3, A->4
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
	// football->1 , basketball->1 , tenis->1

}
