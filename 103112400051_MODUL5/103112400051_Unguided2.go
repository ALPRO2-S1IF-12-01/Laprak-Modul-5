// PRATAMA BINTANG DANISWARA 103112400051
package main

import "fmt"

func bintang(n int) {
	if n <= 0 {
		return
	}
	bintang(n - 1)
	fmt.Println(baris(n))
}

func baris(panjang int) string {
	var hasil string
	for i := 0; i < panjang; i++ {
		hasil += "*"
	}
	return hasil
}

func main() {
	var jumlah int
	fmt.Scan(&jumlah)
	bintang(jumlah)
}
