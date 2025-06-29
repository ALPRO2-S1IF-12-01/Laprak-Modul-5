package main

import "fmt"

func pangkatIteratif(base, exp int) int {
	hasil := 1
	for i := 0; i < exp; i++ {
		hasil *= base
	}
	return hasil
}

func faktorialIteratif(n int) int {
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func main() {
	var base, exp, n int

	fmt.Print("masukkan bilangan: ")
	fmt.Scanln(&base)
	fmt.Print("masukkan pangkat: ")
	fmt.Scanln(&exp)

	fmt.Printf("%d^%d = %d\n", base, exp, pangkatIteratif(base, exp))

	fmt.Print("masukkan angka untuk faktorial: ")
	fmt.Scanln(&n)

	fmt.Printf("%d! = %d\n", n, faktorialIteratif(n))
}
