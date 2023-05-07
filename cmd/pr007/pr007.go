package main

import (
	"fmt"

	"github.com/bcarnazzi/eulergo/pkg/primes"
)

func main() {
	fmt.Println(primes.NthPrime(10_000))
}
