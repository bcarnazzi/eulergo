package main

import (
	"fmt"
	"math/big"

	"github.com/bcarnazzi/eulergo/pkg/primes"
)

func main() {
	n := big.NewInt(600851475143)
	pds := primes.PrimeFactors(n)
	fmt.Println(pds)
	res := big.NewInt(0)
	for _, p := range pds {
		if p.Cmp(res) > 0 {
			res = p
		}
	}
	fmt.Println(res)
}
