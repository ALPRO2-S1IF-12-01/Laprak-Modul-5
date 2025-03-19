//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func cariFaktor(n int, pembagi int) {
    if pembagi > n {
        return
    }
    if n%pembagi == 0 {
        fmt.Printf("%d ", pembagi)
    }
    cariFaktor(n, pembagi+1)
}

func main() {
    var n int

    fmt.Print()
    fmt.Scan(&n)

    if n <= 0 {
        fmt.Println()
        return
    }
    fmt.Print()
    cariFaktor(n, 1)  
    fmt.Println()     
}