# Rangkuman Materi Time Complexity

* Big O adalah sebuah notasi yang digunakan untuk melakukan analisa terhadap sebuah algoritma atau code terhadap waktu eksekusi.

* 1. Constant Time O(1)
```
func dominant (n int) int {
    var result int =0
    result =result +10
    return result
}
```
* 2. Linier Time O(n)
```
funct linier(n int, A[]int) int{
for i:= 0; i <n; i++{
    if A[i] == 0 {
        return 0
    }
}
return 1
}
```
* 3. Linier Time O(n +m)
```
func mergeArrays(arr1 []int, arr2 []int) []int {
    mergedArr := make([]int, 0, len(arr1)+len(arr2))
    
    i := 0
    j := 0
    
    for i < len(arr1) && j < len(arr2) {
        if arr1[i] <= arr2[j] {
            mergedArr = append(mergedArr, arr1[i])
            i++
        } else {
            mergedArr = append(mergedArr, arr2[j])
            j++
        }
    }
    
    for i < len(arr1) {
        mergedArr = append(mergedArr, arr1[i])
        i++
    }
    
    for j < len(arr2) {
        mergedArr = append(mergedArr, arr2[j])
        j++
    }
    
    return mergedArr
}
``` 
* 4. Logarithmic time O(log n)

```
func Logarithmic(n int) int {
    var result int = 0
    for n >1 {
        n/=2
        result +=1
    }
    return result
}
```

* 5. Quadratic Time  O(n^2)

```
func findCommon(arr1 []int, arr2 []int) []int {
    common := make([]int, 0)
    
    for i := 0; i < len(arr1); i++ {
        for j := 0; j < len(arr2); j++ {
            if arr1[i] == arr2[j] {
                common = append(common, arr1[i])
            }
        }
    }
    
    return common
}
```

### Berikut adalah beberapa macam kompleksitas waktu code yang sering ditemui dalam analisis algoritma:

* O(1) - Konstanta
    Kompleksitas waktu O(1) berarti algoritma memiliki waktu eksekusi konstan, tidak tergantung pada ukuran input. Contohnya adalah operasi aritmatika sederhana, seperti penambahan dan pengurangan, atau akses ke elemen array atau variabel.

* O(log n) - Logaritmik
    Kompleksitas waktu O(log n) berarti waktu eksekusi algoritma meningkat secara logaritmik dengan ukuran input. Contohnya adalah algoritma binary search atau pencarian elemen di pohon biner.

* O(n) - Linear
    Kompleksitas waktu O(n) berarti waktu eksekusi algoritma meningkat secara linear dengan ukuran input. Contohnya adalah iterasi pada array atau daftar, atau penjumlahan semua elemen dalam array.

* O(n log n) - Log-linear
    Kompleksitas waktu O(n log n) berarti waktu eksekusi algoritma meningkat dengan faktor logaritmik dari ukuran input. Contohnya adalah algoritma quicksort atau mergesort.

* O(n^2) - Kuadratik
    Kompleksitas waktu O(n^2) berarti waktu eksekusi algoritma meningkat secara kuadratik dengan ukuran input. Contohnya adalah nested loop pada array atau matriks, atau algoritma bubble sort.

* O(2^n) - Exponential
    Kompleksitas waktu O(2^n) berarti waktu eksekusi algoritma meningkat eksponensial dengan ukuran input. Contohnya adalah algoritma brute-force untuk menyelesaikan kombinasi atau subset dari suatu himpunan.

* O(n!) - Faktorial
    Kompleksitas waktu O(n!) berarti waktu eksekusi algoritma meningkat dengan faktorial dari ukuran input. Contohnya adalah algoritma brute-force untuk menyelesaikan masalah permutasi dari suatu himpunan.