package Simple_Unit_Testing

import (
	"fmt"
	"testing"

	"github.com/NaufalFarros/go_Naufal-Farros/23_Unit_Testing/Praktikum/Simple_Unit_Testing"
)

func TestTambah(t *testing.T) {
	hasil := Simple_Unit_Testing.Tambah(2, 3)
	if hasil != 5 {
		t.Errorf("Hasil tambah salah, seharusnya 5, tapi hasilnya %d", hasil)
	}
	fmt.Println("Hasil Test Tambah: ", hasil)
	//tampilkan haisl test
}

func TestKurang(t *testing.T) {
	hasil := Simple_Unit_Testing.Kurang(5, 2)
	if hasil != 3 {
		t.Errorf("Hasil kurang salah, seharusnya 3, tapi hasilnya %d", hasil)
	}
}

func TestBagi(t *testing.T) {
	hasil, err := Simple_Unit_Testing.Bagi(10, 2)
	if err != nil {
		t.Error("Tidak seharusnya terjadi error")
	}
	if hasil != 5 {
		t.Errorf("Hasil bagi salah, seharusnya 5, tapi hasilnya %d", hasil)
	}

	_, err = Simple_Unit_Testing.Bagi(10, 0)
	if err == nil {
		t.Error("Seharusnya terjadi error ketika pembagian dengan angka nol")
	}
	if err.Error() != "Tidak bisa melakukan pembagian dengan angka nol" {
		t.Errorf("Error message salah, seharusnya 'Tidak bisa melakukan pembagian dengan angka nol', tapi messagenya '%s'", err.Error())
	}
}

func TestKali(t *testing.T) {
	hasil := Simple_Unit_Testing.Kali(2, 4)
	if hasil != 8 {
		t.Errorf("Hasil kali salah, seharusnya 8, tapi hasilnya %d", hasil)
	}
}
