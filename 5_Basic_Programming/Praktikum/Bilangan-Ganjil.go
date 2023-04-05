package main

import (
	"fmt"
)

func main() {
  var bilangan int

	fmt.Println("Masukkan Bilangan:")
	fmt.Scan(&bilangan)

	if bilangan%2 == 0 {
		fmt.Println(bilangan, "Bilangan Genap")
	} else {
		fmt.Println(bilangan, "Bilangan Ganjil")
	}
  
}
