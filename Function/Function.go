package main

import "fmt"

func contoh1() {
	fmt.Println("Nama saya adalah Janu")
}

func contoh2(nama string) {
	fmt.Println("Nama saya adalah", nama)
}

func contoh3(a int, b int, c int) int {
	return a + b + c
}

func contoh4(e, f int) (int, int) {
	return e + f, e * f
}

func contoh5(angka ...int) int {
	total := 0
	for _, v := range angka {
		total += v
	}
	return total
}

func main() {
	contoh1()

	contoh2("Cyntia")
	contoh2("Humaira")

	hasil := contoh3(3, 6, 9)
	fmt.Println("Hasil:", hasil)

	jumlah, kali := contoh4(17, 33)
	fmt.Println("Jumlah:", jumlah)
	fmt.Println("Kali:", kali)

	fmt.Println("Total Angka:", contoh5(3, 6, 9))
}
