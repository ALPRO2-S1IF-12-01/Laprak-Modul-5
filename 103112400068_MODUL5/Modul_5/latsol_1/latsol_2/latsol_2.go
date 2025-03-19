package main

import "fmt"

func Star(n int, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	Star(n, i+1)
}
func main() {
	var N int
	fmt.Scan(&N)

	Star(N, 1)
}
