// Anastasia Adinda N.I
// 103112400085
package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%v ", fibonacci(i))
	}
	fmt.Println()
}
