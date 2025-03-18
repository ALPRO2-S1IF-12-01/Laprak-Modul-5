//Muhammad Zaky Mubarok
package main

import "fmt"

func faktor(n int, i int) {
	
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktor(n, i+1)
}

func main() {
	var N int
	fmt.Scanln(&N)

	faktor(N, 1)
	fmt.Println() 
}
