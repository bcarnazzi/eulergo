package fib_test

import (
	"math/big"
	"testing"

	"github.com/bcarnazzi/eulergo/pkg/fib"
)

func Test_MakeFibs(t *testing.T) {
	fibs := fib.MakeFibs(1, 2)

	i := 1
	for v := range fibs {
		if i == 1 && v.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 1, v)
		}
		if i == 2 && v.Cmp(big.NewInt(2)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 2, v)
		}
		if i == 3 && v.Cmp(big.NewInt(3)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 3, v)
		}
		if i == 4 && v.Cmp(big.NewInt(5)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 5, v)
		}
		if i == 5 && v.Cmp(big.NewInt(8)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 8, v)
		}
		if i == 6 && v.Cmp(big.NewInt(13)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 13, v)
		}
		if i == 7 && v.Cmp(big.NewInt(21)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 21, v)
		}
		if i == 8 && v.Cmp(big.NewInt(34)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 34, v)
		}
		if i == 9 && v.Cmp(big.NewInt(55)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 55, v)
		}
		if i == 10 && v.Cmp(big.NewInt(89)) != 0 {
			t.Errorf("incorrect result: expected %d, got %d", 89, v)
		}
		i++
		if i > 10 {
			break
		}
	}
}
