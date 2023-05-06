package primes

import (
	"math/big"
)

//var primeList []*big.Int

func PrimeFactors(n *big.Int) []*big.Int {
	var factors []*big.Int
	nr := new(big.Int).Set(n)

	zero := big.NewInt(0)
	two := big.NewInt(2)

	for new(big.Int).Mod(nr, two).Cmp(zero) == 0 {
		factors = append(factors, big.NewInt(2))
		nr.Div(nr, two)
	}

	i := big.NewInt(3)
	for new(big.Int).Mul(i, i).Cmp(n) <= 0 {
		for new(big.Int).Mod(nr, i).Cmp(zero) == 0 {
			nr.Div(nr, i)
			factors = append(factors, new(big.Int).Set(i))
		}

		i.Add(i, two)
	}

	if nr.Cmp(two) > 0 {
		factors = append(factors, nr)
	}

	return factors
}

func IsPrime(n *big.Int) bool {
	return len(PrimeFactors(n)) == 1
}
