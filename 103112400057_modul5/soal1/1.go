package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku Fibonacci: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Error: Jumlah suku tidak boleh negatif!")
		return
	}

	fmt.Print("Deret Fibonacci: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}
