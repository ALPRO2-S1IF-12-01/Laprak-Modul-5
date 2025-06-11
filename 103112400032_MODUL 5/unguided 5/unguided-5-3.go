// daffa tsaqifna f
package main

import "fmt"

func factorcheck(x, y int) {
	if x%y == 0 {
		fmt.Print(y, " ")
	}
	if x == y {
		return
	}
	factorcheck(x, y+1)
}
func main() {
	var x int
	fmt.Scan(&x)
	factorcheck(x, 1)
}
