## Rangkuman Materi Recursive Sorting dan Searching

* Recursive adalah sebuah langkah algoritma yang mana memanggil fungsi dirinya sendiri

```
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		// cetak perhitungan fibonacci
		// fmt.Println("fibonacci(", n, ") = fibonacci(", n-1, ") + fibonacci(", n-2, ")")
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

```
* Sorting adalah sebuah metode untuk mengurutkan sebuah angka dari yang terkecil maupun terbesar 

```
for i := 0; i < len(productPrice); i++ {
		for j := 0; j < len(productPrice)-1; j++ {
			if productPrice[j] > productPrice[j+1] {
				temp := productPrice[j]
				productPrice[j] = productPrice[j+1]
				productPrice[j+1] = temp
			}
		}
	}
```

* Searching adalah sebuah metode untuk melakukan pencarian 

```
package main

import "fmt"

func linearSearch(arr []int, x int) int {
    for i, val := range arr {
        if val == x {
            return i
        }
    }
    return -1
}

func main() {
    arr := []int{1, 3, 5, 7, 9}
    x := 7
    result := linearSearch(arr, x)
    if result == -1 {
        fmt.Printf("%d not found\n", x)
    } else {
        fmt.Printf("%d found at index,x)
}
```