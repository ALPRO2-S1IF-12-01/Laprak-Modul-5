package main

import (
	"fmt"
)

func printstars(n int) {
	if n == 0 {
		return
	}
	fmt.Print("*")
	printstars(n - 1)
}

func printpola(n, i int) {
	if i > n {
		return
	}
	printstars(i)
	fmt.Println()
	printpola(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	printpola(n, 1)
}
