package main

import "fmt"

func Array() {
	var hari [7]string = [7]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(hari[0])
}

func Slice() {
	var buah = []string{"Apel", "Blueberry", "Cherry"}

	//tambah buah
	buah = append(buah, "Durian")
	fmt.Println("Banyak buah :", len(buah))
}

func Map(){
	var nilai = map[string] int{
		"Janu" : 95,
		"Cyntia" : 80,
	}
	fmt.Println(nilai["Janu"])

	//tambah nilai
	nilai["Huma"] = 85

	//hapus nilai
	delete(nilai, "Cyntia")
}

func SliceMap() {
	var minuman = []string{"kopi", "susu", "jahe", "teh"}
	var harga = map[string] int{
		"kopi" : 8000,
		"susu" : 5000,
		"jahe" : 4000,
		"teh" : 3000,
	}

	var i = ""
	for _, i = range minuman {
		fmt.Println(i, harga[i])
	}
}

func main() {
	Array()
	Slice()
	Map()
	SliceMap()
}