package main

import "fmt"

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""

	//encode chipper
	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			nameEncode += string('a' + (char-'a'+int32(3))%26)
		} else if char >= 'A' && char <= 'Z' {
			nameEncode += string('A' + (char-'A'+int32(3))%26)
		} else {
			nameEncode += string(char)
		}
	}

	return nameEncode

}

func (s *student) Decode() string {
	var nameDecode = ""

	// decode chipper
	for _, char := range s.name {
		if char >= 'a' && char <= 'z' {
			nameDecode += string('a' + (char-'a'+int32(-3))%26)
		} else if char >= 'A' && char <= 'Z' {
			nameDecode += string('A' + (char-'A'+int32(-3))%26)
		} else {
			nameDecode += string(char)
		}
	}

	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student’s name " + a.name + "is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nDecode of student’s name " + a.name + "is : " + c.Decode())
	}
}
