package main

import "fmt"

func main() {
	var s1, s2 int
	for i := 1; i <= 100; i++ {
		s1 += i
	}
	s1 *= s1

	for i := 1; i <= 100; i++ {
		s2 += i * i
	}

	fmt.Println(s1 - s2)
}
