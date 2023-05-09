package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bcarnazzi/eulergo/pkg/files"
)

func calcProduct(s string) uint64 {
	product := uint64(1)
	for i := 0; i < len(s); i++ {
		digit, _ := strconv.Atoi(string(s[i]))
		product *= uint64(digit)
	}
	return product
}

func maxProduct(number string, length int) uint64 {
	maxProduct := uint64(0)
	for i := 0; i <= len(number)-length; i++ {
		//fmt.Println(number[i : i+length])
		product := calcProduct(number[i : i+length])
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct
}

func main() {
	number, err := files.ReadFile(os.Args[1])
	number = strings.Replace(number, "\n", "", -1)
	number = strings.Replace(number, "\r", "", -1)

	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	fmt.Println(maxProduct(number, 13))
}
