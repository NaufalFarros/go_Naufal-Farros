package main

import (
	"fmt"
)

func main() {
//bilangan Prima
	var angka int
	fmt.Println("Masukkan angka:")
	fmt.Scan(&angka)

	for i := 1; i <= angka; i++ {
		var bil = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				bil++
			}
		}

		if bil == 2 {
			fmt.Println(i, "bilangan prima")
		}
	}


}
