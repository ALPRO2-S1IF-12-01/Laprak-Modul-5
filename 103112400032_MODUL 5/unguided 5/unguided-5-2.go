// daffa tsaqifna f
package main

import "fmt"

func lines(x, y int) {
	if x == y {
		return
	}
	star(y)
	lines(x, y+1)
}
func star(x int) {
	if x < 0 {
		fmt.Println()
		return
	}
	fmt.Print("*")
	star(x - 1)
}

func main() {
	var x int
	fmt.Scan(&x)
	lines(x, 0)
}
