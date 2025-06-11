// daffa tsaqifna f
package main

import "fmt"

func back(x, y int) {
	fmt.Print(x)
	if x == y {
		return
	}
	back(x+1, y)
}
func foward(x, y int) {
	fmt.Print(x)
	if x == 1 {
		back(x+1, y)
		return
	}
	foward(x-1, y)
}
func main() {
	var x int
	fmt.Scan(&x)
	foward(x, x)
}
