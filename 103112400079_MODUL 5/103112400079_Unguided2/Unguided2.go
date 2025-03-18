package main

// Muhammad Faris Rachmadi || 103112400079

import "fmt"

func printbintang(n int, baris int) {

	if baris > n {
		return
	}

	for i := 0; i < baris; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	printbintang(n, baris+1)
}

func main() {
	var n int

	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	printbintang(n, 1)
}
