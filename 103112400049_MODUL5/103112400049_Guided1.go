package main

import "fmt"

// fungsi interatif untuk menghitung pangkat (base^exp)
func pangkatIteratif(base, exp int) int {
	hasil := 1

	for i := 0; i < exp; i++ {
		hasil *= base
	}
	return hasil
}

// fungsi interatif untuk menghitung faktorial (n!)
func faktorialIteratif(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}
func main() {
	var base, exp, n int
	// input faktorial
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&base)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scanln(&exp)

	fmt.Printf("%d^%d = %d\n", base, exp, pangkatIteratif(base, exp))
	// input faktorial
	fmt.Print("Masukkan angka untuk faktorial: ")
	fmt.Scanln(&n)

	fmt.Printf("%d! = %d\n", n, faktorialIteratif(n))
}
