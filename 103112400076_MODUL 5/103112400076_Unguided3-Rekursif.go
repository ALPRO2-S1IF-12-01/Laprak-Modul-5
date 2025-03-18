package main

import (
	"fmt"
)

func cetakFaktor(bilangan, pembagi int) {
	if pembagi > bilangan {
		return 
	}

	if bilangan%pembagi == 0 {
		fmt.Printf("%d ", pembagi) 
	}

	cetakFaktor(bilangan, pembagi+1) 
}

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilangan)

	if bilangan < 1 {
		return
	}

	fmt.Printf("Faktor dari %d: ", bilangan)
	cetakFaktor(bilangan, 1) 
	fmt.Println() 
}
