// PRATAMA BINTANG DANISWARA 103112400051
package main

import "fmt"

func faktor(n, i int) {
	if i > n {
		return
	}
	hasil := float64(n) / float64(i)
	if hasil == float64(int(hasil)) {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
	fmt.Println()
}
