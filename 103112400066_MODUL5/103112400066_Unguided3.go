// DWI OKTA SURYANINGRU,
package main

import "fmt"

// Fungsi rekursif untuk mencetak faktor bilangan dari N
func faktor(n, current int) {
	// Basis rekursif: jika current > n, hentikan rekursi
	if current > n {
		return
	}

	// Jika current adalah faktor dari n, cetak current
	if n%current == 0 {
		fmt.Print(current, " ")
	}

	// Rekursif: panggil fungsi faktor untuk nilai current+1
	faktor(n, current+1)
}

func main() {
	var n int

	// Minta pengguna memasukkan nilai n
	fmt.Scan(&n)

	// Panggil fungsi rekursif untuk mencetak faktor bilangan
	faktor(n, 1)
	fmt.Println() // Pindah ke baris baru setelah mencetak faktor
}