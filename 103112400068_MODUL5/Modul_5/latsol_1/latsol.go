package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var jumlahSuku int
	fmt.Scan(&jumlahSuku)

	for i := 0; i <= jumlahSuku; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()
}
