// MUHAMMAD GAMEL AL GHIFARI
//103112400028
package main

import "fmt"

func main() {
    var n int
    fmt.Print("N: ")
    fmt.Scan(&n)
    printPattern(n, 1)
}

func printStars(n int) {
    if n == 0 {
        return
    }
    fmt.Print("*")
    printStars(n - 1)
}

func printPattern(n, current int) {
    if current > n {
        return
    }
    printStars(current)
    fmt.Println()
    printPattern(n, current+1)
}


