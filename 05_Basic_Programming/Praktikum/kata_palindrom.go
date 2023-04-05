package main

import "fmt"

func main() {
  var kata string
	fmt.Print("Masukkan kata : ")
	fmt.Scan(&kata)

	var kataBalik string
	for i := len(kata) - 1; i >= 0; i-- {
		kataBalik += string(kata[i])
	}
	fmt.Println("Captured : ", kataBalik)

	if kata == kataBalik {
		fmt.Println("Kata palindrom")
	} else {
		fmt.Println("Bukan kata palindrom")
	}

}
