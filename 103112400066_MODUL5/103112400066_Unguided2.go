//DWI OKTA SURYANINGRUM
package main

import "fmt"

// Fungsi rekursif untuk mencetak pola bintang
func Bintang(n, current int) {
	// Basis rekursif: jika current > n, hentikan rekursi
	if current > n {
		return
	}

	// Cetak bintang sebanyak current
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println() // Pindah ke baris baru setelah mencetak bintang

	// Rekursif: panggil fungsi Bintang untuk baris berikutnya
	Bintang(n, current+1)
}

func main() {
	var n int

	// Minta pengguna memasukkan nilai n
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	// Panggil fungsi rekursif untuk mencetak pola bintang
	Bintang(n, 1)
}