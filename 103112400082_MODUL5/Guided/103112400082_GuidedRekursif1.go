package main

import "fmt"

// Fungsi Iteratif untuk menghitung pangkat (base^exp)
func PangkatIteratif(base, exp int) int {
	hasil := 1
	for i := 0; i < exp; i++ {
		hasil *= base
	}
	return hasil
}

//Fungsi iteratif untuk menghitung faktorial (n!)
func faktorialIteratif(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil*= i
		
	}
	return hasil
}
func main() {
	var base, exp, n int
	
	// Input Pangkat
	fmt.Print("Masukkan Bilangan: ")
	fmt.Scanln(&base)
	fmt.Print("Masukkan Pangkat: ")
	fmt.Scanln(&exp)

	fmt.Printf("%d^%d = %d\n", base, exp, PangkatIteratif(base, exp))

	//Input faktorial
	fmt.Print("Masukkan angka untuk faktorial: ")
	fmt.Scanln(&n)

	fmt.Printf("%d! = %d\n", n, faktorialIteratif(n))
}