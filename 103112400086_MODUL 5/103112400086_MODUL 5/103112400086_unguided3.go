package main

import "fmt"

func faktorBilangan(n, i int) {
	if i > n { // baris rekursif jika i lebih besar dari n, berenti
		return
	}

	if n%i == 0 { // jika i adalah faktor dari n, cetak
		fmt.Print(i, " ")
	}

	faktorBilangan(n, i+1) // rekursif dengan nilai i bertambah
}

func main() {
	var n int

	fmt.Scanln(&n)

	faktorBilangan(n, 1) // pengecekan dari 1
	fmt.Println()        // baris baru
}
