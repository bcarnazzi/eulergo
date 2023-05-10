package utils

import "math/big"

func IsPalindrome(n int) bool {
	nn := n
	reversed := 0

	for nn > 0 {
		reversed = reversed*10 + nn%10
		nn = nn / 10
	}

	return n == reversed
}

func Gcd(a, b *big.Int) *big.Int {
	return new(big.Int).GCD(nil, nil, a, b)
}

func Lcm(a, b *big.Int) *big.Int {
	p := new(big.Int).Mul(a, b)
	return new(big.Int).Div(p, Gcd(a, b))
}

func IsPythagorean(a, b, c uint64) bool {
	return a < b && b < c && a*a+b*b == c*c
}
