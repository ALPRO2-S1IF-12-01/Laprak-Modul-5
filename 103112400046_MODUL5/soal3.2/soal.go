//ABID FADHILAH MUSTOFA
//103112400046
package main

import "fmt"

func fr(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	fr(n, i+1)
}

func main() {
	var x int
	fmt.Scan(&x)
	fr(x, 1)
	fmt.Println()
}

