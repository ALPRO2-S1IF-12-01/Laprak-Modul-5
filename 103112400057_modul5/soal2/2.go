package main

import "fmt"

func cetakBintang(n int) {
	if n <= 0 {
		fmt.Println("Error: Jumlah baris harus positif!")
		return
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)
	cetakBintang(n)
}
