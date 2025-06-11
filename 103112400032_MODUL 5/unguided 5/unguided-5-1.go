// daffa tsaqifna f
package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Print(Fibonacci(i), " ")
	}
}
