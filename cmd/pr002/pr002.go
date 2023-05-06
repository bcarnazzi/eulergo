package main

import (
	"fmt"
	"math/big"

	"github.com/bcarnazzi/eulergo/pkg/fib"
	"github.com/bcarnazzi/eulergo/pkg/fp"
)

func bigSum(in <-chan *big.Int) *big.Int {
	return fp.Fold(in, big.NewInt(0), func(b1, b2 *big.Int) *big.Int { return new(big.Int).Add(b1, b2) })
}

func main() {
	fibs := fib.MakeFibs(1, 1)
	vb4m := fp.TakeWhile(fibs, func(i *big.Int) bool { return i.Cmp(big.NewInt(4_000_000)) <= 0 })

	fevs := fp.Filter(vb4m, func(b *big.Int) bool {
		var bm big.Int
		bm.Mod(b, big.NewInt(2))
		return bm.Cmp(big.NewInt(1)) == 0
	})

	res := bigSum(fevs)

	fmt.Println(res)

}
