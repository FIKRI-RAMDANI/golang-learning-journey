package main


import "fmt"
// variabel adalah tempat untuk menyimpan data sementara di dalam program
// const adalah tempat untuk menyimpan data yang nilainya tidak bisa diubah

var test = "test"
const alamat = "Jl. Raya No .1"

func main() {
	// variabel dengan type data
	var name string = "Fikri"
	fmt.Println(name)

	// variabel tanpa type data
	var age = 20
	fmt.Println(age)

	// variabel tanpa type data dan tanpa kata kunci var
	country := "Indonesia"
	fmt.Println(country)

	// multi variabel
	var (
		firstName = "Fikri"
		lastName  = "Ramdani"
	)
	fmt.Println(firstName, lastName)

	fmt.Println(alamat)
	fmt.Println(test)
}