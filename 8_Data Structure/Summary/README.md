# Rangkuman Materi Data Structure 

## Array
* Array merupakan sebuah struktur data yang mana dapat diisi dengan berbagai type data. value yang terdapat di dalam array tidak dapat diubah atau berubah
```
namahewan := []string{"harimau", "singa", "gajah", "jerapah"}

```
## Slice 
* Slice merupakan struktur data array yang mana dapat diisi dengan berbagai type data. value yang terdapat di dalam array dapat berubah atau diubah melalui fungsi append.

```
//konversi string ke int
	var angkaInt []int
	for i := 0; i < len(angka); i++ {
		angkaInt = append(angkaInt, int(angka[i]-'0'))
		// fmt.Println(angkaInt)
	}

	//menampilkan angka yang muncul sekali
	var angkaSekali []int
	for i := 0; i < len(angkaInt); i++ {
		var count = 0
		for j := 0; j < len(angkaInt); j++ {
			if angkaInt[i] == angkaInt[j] {
				count++
				// fmt.Println(angkaInt[i], angkaInt[j], count)
			}
		}
		if count == 1 {
			angkaSekali = append(angkaSekali, angkaInt[i])
		}
	}

	return angkaSekali
 ```

 ## Map 
 * Map merupakan struktur data array yang memiliki Key : Primary_Key dan Value. primary key bersifat unik.

 ```
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

``` 






