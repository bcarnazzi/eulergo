package main

import (
	"fmt"
	"math/big"

	"github.com/bcarnazzi/eulergo/pkg/primes"
)

func main() {
	s := big.NewInt(0)
	l := big.NewInt(2_000_000)
	i := 0
	for {
		p := primes.NthPrime(i)
		if p.Cmp(l) < 0 {
			s.Add(s, p)
			i++
		} else {
			break
		}
	}

	fmt.Println(s)
}
