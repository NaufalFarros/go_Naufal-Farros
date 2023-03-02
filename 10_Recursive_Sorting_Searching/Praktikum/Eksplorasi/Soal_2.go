package main

import (
	"fmt"
)

func MaximumBuyProduct(money int, productPrice []int) int {
	for i := 0; i < len(productPrice); i++ {
		for j := 0; j < len(productPrice)-1; j++ {
			if productPrice[j] > productPrice[j+1] {
				temp := productPrice[j]
				productPrice[j] = productPrice[j+1]
				productPrice[j+1] = temp
			}
		}
	}
	// fmt.Println(productPrice)
	productbuy := 0

	for i := 0; i < len(productPrice); i++ {
		if money >= productPrice[i] {
			money -= productPrice[i]
			productbuy++
		}
	}

	return productbuy
}

func main() {
	fmt.Println(MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))      // 3
	fmt.Println(MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})) // 4
	fmt.Println(MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}))   // 4
	fmt.Println(MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000}))           // 1
	fmt.Println(MaximumBuyProduct(0, []int{10000, 30000}))                        // 0
}
