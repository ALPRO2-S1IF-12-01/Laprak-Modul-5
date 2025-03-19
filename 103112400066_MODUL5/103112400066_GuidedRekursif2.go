//  DWI OKTA SURYANINGRUM
package main

import "fmt"

// FUNGSI REKURSIF UNTUK MENGHITUNG PANGKAT (BASE^EXP)
func pangkatRekursif(base, exp int) int {
	// Jika exp (pangkat) adalah 0, kembalikan 1, Karena setiap bilangan yang dipangkatkan 0 hasilnya adalah 1
	if exp == 0 {
		return 1
	}
	// Mengembalikan base dikali dengan pangkatRekursif(base, exp-1)
	// Ini akan terus mengurangi exp hingga mencapai 0
	return base * pangkatRekursif(base, exp-1)
}

// FUNGSI REKURSIF UNTUK MENGHITUNG FAKTORIAL
func faktorialRekursif(n int) int {
	// Jika n adalah 0 atau 1, kembalikan 1, Karena faktorial dari 0 dan 1 adalah 1
	if n == 0 || n == 1 {
		return 1
	}
	// mengembalikan n dikali dengan faktorialRekursif(n-1)
	// Ini akan terus mengurangi n hingga mencapai 0 atau 1
	return n * faktorialRekursif(n-1)
}

func main() {
	// Variabel untuk menyimpan input dari pengguna
	var base, exp, n int

	// INPUT PANGKAT
	// Minta pengguna memasukkan bilangan (base) dan pangkat (exp)
	fmt.Print("Masukkan Bilangan : ")
	fmt.Scanln(&base)
	fmt.Print("Masukkan Pangkat : ")
	fmt.Scanln(&exp)

	// Cetak hasil perhitungan pangkat menggunakan fungsi pangkatRekursif
	fmt.Printf("%d^%d = %d\n", base, exp, pangkatRekursif(base, exp))

	// INPUT FAKTORIAL
	// Minta pengguna memasukkan angka untuk menghitung faktorial
	fmt.Print("Masukkan Angka Untuk Faktorial : ")
	fmt.Scanln(&n)

	// Cetak hasil perhitungan faktorial menggunakan fungsi faktorialRekursif
	fmt.Printf("%d! = %d\n", n, faktorialRekursif(n))
}