package main

import "fmt"

func pangkatRekursif(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * pangkatRekursif(base, exp-1)
}

func faktorialrekursif(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorialrekursif(n-1)
}

func main() {
	var base, exp, n int
	fmt.Print("Masukkan  bilangan: ")
	fmt.Scanln(&base)
	fmt.Print("masukkan pangkat: ")
	fmt.Scanln(&exp)

	fmt.Printf("%d^%d = %d\n", base, exp, pangkatRekursif(base, exp))
	fmt.Scanln(&n)
	fmt.Printf("%d! = %d\n", n, faktorialrekursif(n))
}
