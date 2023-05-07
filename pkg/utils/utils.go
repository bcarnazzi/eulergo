package utils

func IsPalindrome(n int) bool {
	nn := n
	reversed := 0

	for nn > 0 {
		reversed = reversed*10 + nn%10
		nn = nn / 10
	}

	return n == reversed
}
