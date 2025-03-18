package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	var n int
	fmt.Print("Masukkan jumlah suku Fibonacci: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Masukkan angka positif!")
		return
	}

	fmt.Println("Deret Fibonacci:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}
