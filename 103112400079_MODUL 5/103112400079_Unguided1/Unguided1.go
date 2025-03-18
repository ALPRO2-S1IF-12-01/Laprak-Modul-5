package main

// Muhammad Faris Rachmadi || 103112400079

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println("Deret Fibonacci hingga ke-", n)
	for i := 0; i <= n; i++ {
		fmt.Printf("Suku ke-%d: %d\n", i, fibonacci(i))
	}
}
