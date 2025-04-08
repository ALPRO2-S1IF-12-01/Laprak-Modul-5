package main

import "fmt"

func asteriks(n, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	asteriks(n, i + 1)
}

func main() {
	var n int
	fmt.Scan(&n)

	asteriks(n, 1)
}
