package main

import (
	"fmt"
)

func munculSekali(angka string) []int {
	//konversi string ke int
	var angkaInt []int
	for i := 0; i < len(angka); i++ {
		angkaInt = append(angkaInt, int(angka[i]-'0'))
		// fmt.Println(angkaInt)
	}

	//menampilkan angka yang muncul sekali
	var angkaSekali []int
	for i := 0; i < len(angkaInt); i++ {
		var count = 0
		for j := 0; j < len(angkaInt); j++ {
			if angkaInt[i] == angkaInt[j] {
				count++
				// fmt.Println(angkaInt[i], angkaInt[j], count)
			}
		}
		if count == 1 {
			angkaSekali = append(angkaSekali, angkaInt[i])
		}
	}

	return angkaSekali

}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
