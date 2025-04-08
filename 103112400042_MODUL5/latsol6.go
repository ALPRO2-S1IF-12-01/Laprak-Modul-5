// M.HANIF AL FAIZ
// 103112400042
package main

import (
	"fmt"
)

func pangkat(x int, y int) float64 {
	if y < 0 {
		return 1 / pangkat(x, -y)
	} else if y == 0 {
		return 1
	} else if x == 0 {
		return 0
	} else {
		return float64(x) * pangkat(x, y-1)
	}
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan bulat x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat y: ")
	fmt.Scan(&y)

	hasil := pangkat(x, y)
	fmt.Printf("%d dipangkatkan %d = %.2f\n", x, y, hasil)
}
