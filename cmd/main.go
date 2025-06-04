package main

import (
	"fmt"

	"github.com/popescuag/euler/even-fibo/evenfibo"
)

func main() {
	fmt.Printf("Sum of even Fibonacci numbers up to 4 million: %d\n", evenfibo.EvenFibonacci(4000000))
}
