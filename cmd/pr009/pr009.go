package main

import (
	"fmt"
	"os"

	"github.com/bcarnazzi/eulergo/pkg/utils"
)

func main() {
	for a := uint64(1); a < 1000; a++ {
		for b := a; b < 1000; b++ {
			if a+b < 1000 {
				c := 1000 - a - b
				if utils.IsPythagorean(a, b, c) {
					fmt.Printf("%d * %d * %d = %d\n", a, b, c, a*b*c)
					os.Exit(0)
				}
			}
		}
	}
}
