package fib

import (
	"math/big"

	"github.com/bcarnazzi/eulergo/pkg/fp"
)

func MakeFibs(u0, u1 int64) chan *big.Int {
	return fp.Unfold2(big.NewInt(u0), big.NewInt(u1),
		func(i1, i2 *big.Int) *big.Int { return new(big.Int).Add(i1, i2) })
}
