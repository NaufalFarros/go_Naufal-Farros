package simple_unit_testing

import (
	"errors"
)

func Tambah(angka1 int, angka2 int) int {
	return angka1 + angka2
}

func Kurang(angka1 int, angka2 int) int {
	return angka1 - angka2
}

func Bagi(angka1 int, angka2 int) (int, error) {
	if angka2 == 0 {
		return 0, errors.New("Tidak bisa melakukan pembagian dengan angka nol")
	}
	return angka1 / angka2, nil
}

func Kali(angka1 int, angka2 int) int {
	return angka1 * angka2
}
