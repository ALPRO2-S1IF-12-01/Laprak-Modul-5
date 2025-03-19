// Nama: Anggun Wahyu Widiyana (103112480280)
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 { // Base-case pertama
		return 0
	} else if n == 1 { // Base-case kedua
		return 1
	} else { // Recursive-case
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	var n int
	fmt.Scan(&n) 

	fmt.Printf("Nilai suku ke-%d dari deret Fibonacci adalah %d\n", n, fibonacci(n))
}