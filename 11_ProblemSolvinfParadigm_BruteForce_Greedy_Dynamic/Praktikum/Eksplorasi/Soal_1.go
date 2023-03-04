package main

import "fmt"

func toRoman(num int) string {
	if num < 1 || num > 3999 {
		return "Invalid input"
	}

	nilaiRomawi := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	simbolRomawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i := 0; i < len(nilaiRomawi); i++ {
		for num >= nilaiRomawi[i] {
			result += simbolRomawi[i]
			num -= nilaiRomawi[i]
		}
	}

	return result
}

func main() {
	fmt.Println(toRoman(4))    //IV
	fmt.Println(toRoman(9))    //IX
	fmt.Println(toRoman(23))   //XXIII
	fmt.Println(toRoman(2021)) //MMXXI
	fmt.Println(toRoman(1646)) //MDCXLVI
}
