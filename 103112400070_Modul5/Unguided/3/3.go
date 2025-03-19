//Achmad Zulvan Nur Hakim 103112400070
package main

import 	"fmt"

func faktor(x, i int) {
	if i > x {
		return
	}
	if x%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(x, i+1)
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Print("Faktor ", x, " : ")
	faktor(x, 1)
}
