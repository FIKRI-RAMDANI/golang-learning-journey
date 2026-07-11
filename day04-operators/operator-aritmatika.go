package main

import "fmt"

func main() {
	var a = 10
	var b = 11
	var c = 20

// pertambahan
	var d = a + b + c
	fmt.Println("hasil dari a + b + c =", d)

// pengurangan
	var e = c - b
	fmt.Println(e)

// pembagaian
	var g = c / b
	fmt.Println("Hasil dari c / b =", g)

// perkalian
	var h = a * b
	fmt.Println("Hasil dari a * b =", h)

// Pembagian Integer dengan Float
	var i = 10 / 3
	fmt.Println(i)  // hasilnya 3, bukan 3.333 (karena int dibagi int = intiger, desimal dibuang)

	var j = 10.0 / 3.0
	fmt.Println(j) // hasilnya 3.333333 (karena menggunakan float)

// modulus / sisa pembagi
	n := 20
	if n % 2 == 0 {
		fmt.Println("genap")
	} else {
		fmt.Println("ganjil")
	}

// compound assigment operator (+=, -=, *=, /=, %=)
	m := 20
	m += 5 // sama dengan m = m + 5
	fmt.Println(m)
}