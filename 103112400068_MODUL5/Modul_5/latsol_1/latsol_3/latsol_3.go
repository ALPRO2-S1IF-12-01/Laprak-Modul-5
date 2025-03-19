package main

import "fmt"

func faktorRekursif(N, i int) {
	if i > N {
		return
	}
	if N%i == 0 {
		fmt.Printf("%d ", i)
	}
	faktorRekursif(N, i+1)
}

func main() {
	var N int
	fmt.Scanln(&N)
	fmt.Print("Faktor: ")
	faktorRekursif(N, 1)
	fmt.Println()
}
