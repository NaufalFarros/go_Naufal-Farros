package main

import (
	"fmt"
)

type Car struct {
	Type   string
	FuelIn float32
}

func (car Car) EstimasiJarak() float32 {

	//deklarasi variabel
	var FuelInMililiter float32 = car.FuelIn

	//menghtung jarak yang bisa ditempuh
	var jarak = FuelInMililiter / 1.5

	return jarak

}

func main() {
	car := Car{
		Type:   "Sedan",
		FuelIn: 10,
	}

	jarakYangDitempuh := car.EstimasiJarak()
	//konversi mililiter ke kilometer

	fmt.Println("estimasi jarak yang ditempuh", jarakYangDitempuh, "mill")

}
