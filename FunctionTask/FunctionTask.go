/*
1. Buat fungsi bernama kategorikanBilangan(n int) string
2. Fungsi ini akan:
- Mengembalikan "Kelipatan Enam" jika n % 6 == 0
- "Genap" jika n % 2 == 0
- "Ganjil" jika tidak
3. Gunakan fungsi itu di dalam for loop dari angka 1–30 dan cetak hasil seperti ini:
*/

package main

import "fmt"

func kategorikanBilangan(n int) string {
	if n%6 == 0 {
		return "Kelipatan Enam"
	} else if n%2 == 0 {
		return "Genap"
	} else {
		return "Ganjil"
	}
}

func main() {
	for n := 1; n <= 30; n++ {
		fmt.Println(n, kategorikanBilangan(n))
	}
}
