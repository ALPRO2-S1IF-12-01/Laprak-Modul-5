package main

//SETYO NUGROHO
//103112400024
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cetakBintang(n, 1)
}

func cetakBintang(n, i int) {
	if i > n {
		return
	}
	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()
	cetakBintang(n, i+1)
}
