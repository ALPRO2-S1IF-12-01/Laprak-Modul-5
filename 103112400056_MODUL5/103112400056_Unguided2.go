// Felix Pedrosa V 

package main

import "fmt"

// Fungsi rekursif untuk menampilkan bintang
func printStars(n int, current int) {
	if current > n {
		return
	}
	// Mencetak bintang sesuai dengan angka current
	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println() // Pindah ke baris berikutnya
	printStars(n, current+1) // Panggil fungsi rekursif untuk current yang berikutnya
}

func main() {
	var N int
	fmt.Scan(&N) // Membaca input dari user

	printStars(N, 1) // Memanggil fungsi untuk mulai dari 1
}