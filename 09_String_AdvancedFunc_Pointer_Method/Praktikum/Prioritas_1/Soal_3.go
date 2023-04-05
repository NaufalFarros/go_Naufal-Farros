package main

import (
	"fmt"
	"strings"
)

func compare(a, b string) string {

	// cek apakah ada kata yang sama
	if strings.Contains(a, b) {
		return b
	}

	return a

}

func main() {
	fmt.Println(compare("AKA", "AKASHI"))     //AKA
	fmt.Println(compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(compare("KI", "KIJANG"))      //KI
	fmt.Println(compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(compare("ILALANG", "ILA"))    //ILA

}
