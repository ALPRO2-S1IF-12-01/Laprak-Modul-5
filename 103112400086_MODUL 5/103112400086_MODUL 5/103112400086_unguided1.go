package main

import "fmt"

// Fungsi rekursif untuk menghitung Fibonacci
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

	fmt.Print("n : ")
	fmt.Scanln(&n)

	fmt.Print("Sn: ")
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i))
		if i < n {
			fmt.Print(" ")
		}
	}
}
