// Raihan Adi Arba
// 103112400071
package main

import "fmt"

func deretFibonacci(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return deretFibonacci(x-1) + deretFibonacci(x-2)
}

func main() {
	var n int
	fmt.Scan(&n)

	// Cetak deret Fibonacci
	for i := 0; i <= n; i++ {
		fmt.Print(deretFibonacci(i))
		if i < n {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
