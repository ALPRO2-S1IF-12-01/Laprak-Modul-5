//MUHAMMAD GAMEL AL GHIFARI
//103112400028
package main
import "fmt"
func main() {
    var n int
    fmt.Print("Masukkan n: ")
    fmt.Scan(&n)

    fmt.Printf("Faktor dari %d: ", n)
    printFactors(n, 1)
    fmt.Println()
}

func printFactors(n, i int) {
    if i > n {
        return
    }
    if n%i == 0 {
        fmt.Printf("%d ", i)
    }
    printFactors(n, i+1)
}


