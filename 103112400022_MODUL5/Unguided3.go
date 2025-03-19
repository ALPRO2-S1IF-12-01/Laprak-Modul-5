//Damanik, Yohanes Geovan Ondova
//103112400022

package main  
import "fmt"

func faktorRekursif(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorRekursif(n, i+1)
}

func main(){
	var N int
	fmt.Scan(&N)
	faktorRekursif(N, 1)
	fmt.Println()
}