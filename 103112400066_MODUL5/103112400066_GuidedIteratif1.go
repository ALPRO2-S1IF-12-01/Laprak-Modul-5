// DWI OKTA SURYANINGRUM
package main

import "fmt"

// FUNGSI ITERATIF UNTUK MENGHITUNG PANGKAT (BASE^EXP)
// Fungsi ini digunakan untuk menghitung hasil dari base yang dipangkatkan dengan exp secara iteratif (dengan perulangan).
// Mulai dengan hasil = 1, kemudian kalikan hasil dengan base sebanyak exp kali.
func pangkatIteratif(base, exp int) int {
	// Inisialisasi hasil dengan 1
	hasil := 1
	// Perulangan sebanyak exp kali
	for i := 0; i < exp; i++ {
		// Kalikan hasil dengan base setiap kali perulangan
		hasil *= base
	}
	// Kembalikan hasil akhir setelah perulangan selesai
	return hasil
}

// FUNGSI ITERATIF UNTUK MENGHITUNG FAKTORIAL
// Fungsi ini menghitung faktorial dari angka n secara iteratif.
// Mulai dengan hasil = 1, kemudian kalikan hasil dengan angka dari 2 hingga n.
func faktorialIteratif(n int) int {
	// Inisialisasi hasil dengan 1
	hasil := 1
	// Perulangan dari 2 hingga n
	for i := 2; i <= n; i++ {
		// Kalikan hasil dengan i setiap kali perulangan
		hasil *= i
	}
	// Kembalikan hasil faktorial setelah perulangan selesai
	return hasil
}

func main() {
	// Mendeklarasikan variabel untuk menampung input base, exp, dan n
	var base, exp, n int

	// INPUT UNTUK MENGHITUNG PANGKAT
	// Meminta pengguna untuk memasukkan angka 'base' (bilangan dasar) dan 'exp' (pangkat)
	fmt.Print("Masukkan Bilangan : ")
	fmt.Scanln(&base)  // Membaca input untuk base
	fmt.Print("Masukkan Pangkat : ")
	fmt.Scanln(&exp)   // Membaca input untuk exp (pangkat)

	// Menampilkan hasil pangkat (base^exp) dengan memanggil fungsi 'pangkatIteratif'
	fmt.Printf("%d^%d = %d\n", base, exp, pangkatIteratif(base, exp))

	// INPUT UNTUK MENGHITUNG FAKTORIAL
	// Meminta pengguna untuk memasukkan angka 'n' untuk dihitung faktorialnya
	fmt.Print("Masukkan Angka Untuk Faktorial : ")
	fmt.Scanln(&n) // Membaca input untuk 'n'

	// Menampilkan hasil faktorial dari angka 'n' dengan memanggil fungsi 'faktorialIteratif'
	fmt.Printf("%d! = %d\n", n, faktorialIteratif(n))
}