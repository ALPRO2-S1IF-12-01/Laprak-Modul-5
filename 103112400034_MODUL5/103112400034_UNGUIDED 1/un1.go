// MULIA AKBAR NANDA PRATAMA
// 103112400034
package main

import "fmt"

func fibonacci(n int) int {
        if n <= 1 {
                return n
        }
        return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("n: ")
	fmt.Scanln(&n)

	fmt.Print("Sn: ")
        for i := 0; i <= 10; i++ {
                fmt.Printf("%d ", fibonacci(i))
        }
        fmt.Println()
}