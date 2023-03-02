package main

import "fmt"

func caesar(offset int, input string) string {
	var result string
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			idx := (int(char-'a') + offset) % len(alphabet)
			result += string(alphabet[idx])
		} else if char >= 'A' && char <= 'Z' {
			idx := (int(char-'A') + offset) % len(alphabet)
			result += string(alphabet[idx] - 32) // ubah ke huruf besar
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
