// DWI OKTA SURYANINGRUM
package main

import "fmt"

// Fungsi rekursif untuk menghitung nilai deret Fibonacci ke-n
func fibonacci(n int) int {
	// Jika n adalah 0 atau 1, kembalikan nilai n
	if n == 0 || n == 1 {
		return n
	}
	// Mengembalikan penjumlahan dua suku sebelumnya
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int

	// Meminta pengguna memasukkan nilai n
	fmt.Print("Masukkan nilai n (suku ke-n): ")
	fmt.Scan(&n)

	// Cetak nilai deret Fibonacci dari suku ke-0 hingga suku ke-n
	fmt.Println("Deret Fibonacci hingga suku ke-", n, ":")
	for i := 0; i <= n; i++ {
		fmt.Println(fibonacci(i))
	}
}