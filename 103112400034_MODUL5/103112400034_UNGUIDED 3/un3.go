// MULIA AKBAR NANDA PRATAMA
// 103112400034
package main

import "fmt"
func cariFaktor(n int, faktor int) {
        if faktor > n {
                return
        }
        if n%faktor == 0 {
                fmt.Printf("%d ", faktor)
        }
        cariFaktor(n, faktor+1)
}

func main() {
        var n int
        fmt.Print("N: ")
        fmt.Scan(&n)

        fmt.Printf("Faktor dari %d adalah: ", n)
        cariFaktor(n, 1)
        fmt.Println()
}