package main

import "fmt"

func pangkatIteratif(base, exp int) int{
	hasil := 1
	for i := 0; i < exp; i++ {
		hasil *= base
	}
	return hasil
}

func faktorialIteratif(n int) int{
	hasil := 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func main() {
	var base, exp, n int
	//Input pangkat
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&base)

	fmt.Print("Masukkan pangkat: ")
	fmt.Scan(&exp)

	fmt.Printf("%d^%d = %d\n", base, exp, pangkatIteratif(base, exp))

	//input faktorial
	fmt.Print("Masukkan angka untuk faktorial: ")
	fmt.Scan(&n)

	fmt.Printf("%d! = %d\n", n,  faktorialIteratif(n))

}
