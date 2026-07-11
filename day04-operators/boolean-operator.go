package main

import "fmt"

func main() {
	// operators AND && (kedua kondisi harus true/benar)
	umur := 20
	punyaKtp := true
	if umur == 20 && punyaKtp {
		fmt.Println("Sudah Legal")
	} else {
		fmt.Println("Belum Legal")
	}

	// Operator OR || (salah satu kondisi benar)
	if umur < 20 || punyaKtp {
		fmt.Println("Sudah Legal")
	} else {
		fmt.Println("Belum Legal")
	}

	// Operator NOT/kebalikan ! (membalikan nilai boolean)
	isActive := false

	if !isActive {
		fmt.Println("User tidak aktif")
	}
}