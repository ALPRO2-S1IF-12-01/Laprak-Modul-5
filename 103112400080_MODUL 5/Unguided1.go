// JESIKA METANIA RAHMA ARIFIN
//103112400080
package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println("n\t", "Sₙ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%d\t %d\n", i, fibonacci(i))
	}
}
