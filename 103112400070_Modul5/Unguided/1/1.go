// Achmad Zulvan Nur Hakim 103112400070
package main

import 	"fmt"


func deret(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return deret(n-1) + deret(n-2)
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Print("Deret :")
	for i := 0; i <= x; i++ {
		fmt.Print(deret(i), " ")
	}
}
