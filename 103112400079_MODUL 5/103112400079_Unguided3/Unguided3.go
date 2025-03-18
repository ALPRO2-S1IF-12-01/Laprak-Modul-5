package main

//Muhammad Faris Rachmadi||103112400079
import "fmt"

func carifaktor(n int, i int) {

	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Print(i, " ")
	}

	carifaktor(n, i+1)
}

func main() {
	var n int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	carifaktor(n, 1)
	fmt.Println()
}
