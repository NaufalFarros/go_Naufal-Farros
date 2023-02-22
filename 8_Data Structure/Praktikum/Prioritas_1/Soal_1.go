package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {

	var c []string
	for i := 0; i < len(arrayA); i++ {
		c = append(c, arrayA[i])
	}
	for i := 0; i < len(arrayB); i++ {
		c = append(c, arrayB[i])
	}

	//jika ada yang sama maka hapus
	for i := 0; i < len(c); i++ {
		for j := i + 1; j < len(c); j++ {
			if c[i] == c[j] {
				c = append(c[:j], c[j+1:]...)

			}
		}
	}

	//berikan pemisah tanda petik dan koma
	for i := 0; i < len(c); i++ {
		c[i] = "\"" + c[i] + "\"" + ","
	}

	return c

}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"addie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "addie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
