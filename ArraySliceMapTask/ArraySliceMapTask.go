/* 
Buat program yang:
1. Menyimpan daftar 5 **nama mahasiswa** dalam slice.
2. Menyimpan **nilai** tiap mahasiswa dalam map (`map[string]int`).
3. Loop seluruh mahasiswa dan tampilkan:
    - Nama
    - Nilai
    - Keterangan kelulusan:
        - Nilai ≥ 75 → "Lulus"
        - < 75 → "Tidak Lulus"
*/

package main

import "fmt"

func main() {

	var nama = []string{"Janu", "Rizqi", "Dwi", "Cyntia", "Huma"}
	var nilai = map[string]int{
		"Janu":   100,
		"Rizqi":  90,
		"Dwi":    80,
		"Cyntia": 70,
		"Huma":   60,
	}

	for _, n := range nama {
		score := nilai[n]
		status := "Tidak Lulus"
		if score >= 75 {
			status = "Lulus"
		}
		fmt.Printf("Nama: %s, Nilai: %d, Keterangan: %s\n", n, score, status)
	}
}
