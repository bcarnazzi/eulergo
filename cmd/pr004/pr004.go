package main

import (
	"fmt"

	"github.com/bcarnazzi/eulergo/pkg/utils"
)

func main() {
	lp := 0
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			p := i * j
			if utils.IsPalindrome(p) && p > lp {
				lp = p
			}
		}
	}
	fmt.Println(lp)
}
