package main

import (
	"fmt"
	"math/big"

	"github.com/bcarnazzi/eulergo/pkg/utils"
)

func main() {
	twenty := big.NewInt(20)
	res := big.NewInt(1)

	for i := big.NewInt(1); i.Cmp(twenty) <= 0; i.Add(i, big.NewInt(1)) {
		res = utils.Lcm(res, i)
	}

	fmt.Println(res)
}
