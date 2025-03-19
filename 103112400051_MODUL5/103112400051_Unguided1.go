// PRATAMA BINTANG DANISWARA 103112400051
package main

import "fmt"

func fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}
