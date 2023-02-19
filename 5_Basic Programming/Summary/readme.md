## Rangkuman Materi Basic Program
1. Mempelajari dasar bahasa pemrograman go
   * Go merupakan bahasa merograman yang mirip seperti bahasa merograman java dan C++, Go dibuat oleh Robert Grisemer, Rob Pike, Ken Thompson, lan Lance Taylor, dan Russ Cox. bahasa pemrograman go pertama kali rislis pada bulan november tahun 2009
2. Memahami type data 
   * Terdapat beberapa type data pada bahasa pemrograman Go yaoitu Boolean, Numeric, dan String
3. Memahami perbedaan Const dan Var
   * Const merupakan variabel yang tidak dapat diperbarui dan dideklarasi ulang
   * Var meruapakan variabel scope global yang dapat diperbarui dan dideklarasi ulang
4. Memahami Operator di Golang
   * Operator Addition (+), Subtraction (-), Division (/), Multiplication, Modulo (%), Increment (++), Decrement (--)
5. Memahami konsep Pengkondisian suatu persoalan
   * Contoh dalam menentukan bilangan ganjil dan genap
``` go
  angka := 4
  
  if angka %2 == 0 {
    fmt.Println("Genap")
  }else {
    fmt.Println("Ganjil")
  }
```
6. Memahami konsep perulangan pada pemrograman golang
   * Contoh Perulangan Looping
 ``` go
 for i:=1; i<=15; i++ {
  fmt.Println(i)
 }
```
7. Memahami konsep nested loop
 * Contoh Nested Loop
``` go
for i:=1; i<=15; i++ {
   for i:=1; i<=5; i++ {
   fmt.Println("Hallo")
  }
}
```
