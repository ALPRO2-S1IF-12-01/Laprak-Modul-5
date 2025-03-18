// Felix Pedrosa V 

package main

import "fmt"

// Fungsi rekursif untuk mencetak faktor
func faktorRekursif(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorRekursif(n, i+1)
}

func main() {
	var N int

	// Membaca input dari pengguna
	fmt.Scan(&N)

	faktorRekursif(N, 1) // Memanggil fungsi rekursif dimulai dari 1
	fmt.Println() // Untuk newline setelah output
}