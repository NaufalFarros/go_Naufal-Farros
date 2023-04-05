package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

func (s Student) NilaiTertinggi(students []Student) (string, int) {

	max := 0
	var nama string
	for _, s := range students {
		if s.Score > max {
			max = s.Score
			nama = s.Name
		}
	}
	return nama, max
}

func (s Student) NilaiTerendah(students []Student) (string, int) {
	min := math.MaxInt32
	var nama string
	for _, s := range students {
		if s.Score < min {
			min = s.Score
			nama = s.Name
		}
	}
	return nama, min
}

func main() {
	var name string
	var score int
	var studentSlice []Student

	for i := 1; i <= 5; i++ {
		fmt.Print("Masukkan nama ke- ", i, " :")
		fmt.Scan(&name)
		fmt.Print("Masukkan nilai ke-", i, " :")
		fmt.Scan(&score)
		students := Student{
			Name:  name,
			Score: score,
		}
		studentSlice = append(studentSlice, students)
	}

	// Mencari nilai tertinggi dari slice studentSlice
	namaNilaiTertinggi, nilaiTertinggi := studentSlice[0].NilaiTertinggi(studentSlice)
	namaNilaiTerendah, nilaiTerendah := studentSlice[0].NilaiTerendah(studentSlice)

	var jumlahNilai int = 0
	for _, data := range studentSlice {
		jumlahNilai = jumlahNilai + data.Score

		// fmt.Println("Jumlah :", jumlahNilai)
	}
	var rata_rata = jumlahNilai / len(studentSlice)

	fmt.Println("Rata-rata nilai siswa :", rata_rata)
	fmt.Println("Nilai tertinggi siswa :", namaNilaiTertinggi, "(", nilaiTertinggi, ")")
	fmt.Println("Nilai terendah siswa :", namaNilaiTerendah, "(", nilaiTerendah, ")")

}
