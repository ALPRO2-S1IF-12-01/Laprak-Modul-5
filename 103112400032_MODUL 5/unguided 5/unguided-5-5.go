// daffa tsaqifna f
package main

import "fmt"

func ganjil(x int, y int) {
	if y > x {
		return
	}
	if y%2 == 1 {
		fmt.Printf("%d ", y)
	}
	ganjil(x, y+1)
}

func main() {
	var x int
	fmt.Scan(&x)
	ganjil(x, 1)
}
