/* Buat program Go untuk:
1. Menampilkan bilangan dari 1 sampai 30.
2. Jika bilangan ganjil → tampilkan "Ganjil"
3. Jika bilangan genap → tampilkan "Genap"
4. Jika bilangan kelipatan 6 → tampilkan "Kelipatan Enam"
5. Gunakan if dan for secara tepat. */

package main

import "fmt"

func main() {
	for i := 1; i <= 30; i++ {
		if i%6 == 0 {
			fmt.Println(i, ": Kelipatan Enam")
		} else if i%2 == 0 {
			fmt.Println(i, ": Genap")
		} else {
			fmt.Println(i, ": Ganjil")
		}
	}
}
