package main

import (
	"fmt"
)

func main() {
  var alas, tinggi, sisiAtas float64

	fmt.Println("Masukkan alas:")
	fmt.Scan(&alas)

	fmt.Println("Masukkan tinggi:")
	fmt.Scan(&tinggi)

	fmt.Println("Masukkan sisiAtas:")
	fmt.Scan(&sisiAtas)

	var luas float64 = (alas + sisiAtas) * tinggi / 2

	fmt.Println("Luas trapesium adalah:", luas)


}
