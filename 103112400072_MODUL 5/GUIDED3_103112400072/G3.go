package main

import (
	"fmt"
)

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
	fmt.Print("Masukan bilangan : ")
	fmt.Scanln(&base)

	fmt.Print("Masukan Pangkat : ")
	fmt.Scanln(&exp)
    
	fmt.Printf("%d^%d\n", base, exp, pangkatIteratif(base, exp))

	fmt.Print("Masukan Faktorial : ")
	fmt.Scan(&n)

	fmt.Printf("%d! = %d\n", n,faktorialIteratif(n))
}