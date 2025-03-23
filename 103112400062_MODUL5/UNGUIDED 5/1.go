// M.DAVI ILYAS RENALDO
// 103112400062
package main

import "fmt"

func bintang(n, x int) {
	if x > n {
		return
	}
	for i := 0; i < x; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	bintang(n, x+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	bintang(n, 1)
}