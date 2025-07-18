package main

import "fmt"

func main() {
	nama := "Januarizqi Dwi Milleniantoro"
	umur := 25
	var tinggi float64 = 1.58
	var berat float64 = 57
	isBelajar := true
	belajar := ""

	var bmi float64 = 0
	var kategori string = ""
	bmi = berat / (tinggi * tinggi)

	if isBelajar == true {
		belajar = "Benar"
	} else {
		belajar = "Salah"
	}

	if bmi < 18.5 {
		kategori = "Kurus"
	} else if bmi <= 24.9 {
		kategori = "Normal"
	} else if bmi <= 29.9 {
		kategori = "Gemuk"
	} else {
		kategori = "Obesitas"
	}

	fmt.Println("Nama saya adalah", nama)
	fmt.Println("Umur saya sekarang adalah", umur)
	fmt.Println("Tinggi saya sekarang adalah", tinggi)
	fmt.Println("Berat saya sekarang adalah", berat)
	fmt.Println("Saya belajar Go:", belajar)

	fmt.Printf("\nBMI: %.2f", bmi)
	fmt.Println("\nKategori:", kategori)

}
