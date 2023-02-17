package main

import (
	"fmt"
)

func main() {
//Bialngan kelipatan 7
	var bilangan int

	fmt.Println("Masukkan Bilangan:")
	fmt.Scan(&bilangan)

	if bilangan%7 == 0 {
		fmt.Println(bilangan, "Bilangan kelipatan 7")
	} else {
		fmt.Println(bilangan, "Bilangan bukan kelipatan 7")

	}

}
