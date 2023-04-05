package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
	var m = make(map[string]int)
	for _, v := range slice {
		m[v]++
	}
	return m
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) // map[asd:2 qwe:3 adi:1]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      // map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         // map[]
}
