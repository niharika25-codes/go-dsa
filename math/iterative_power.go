package main

import "fmt"

func power(x, y int) int {
	res := 1
	for y > 0 {
		if y%2 != 0 {
			res *= x
		}
		x *= x
		y /= 2
	}
	return res
}

func powerIterative(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		x = 1 / x
		n = -n
	}

	for n > 0 {
		// Bitwise check: n&1 is a faster way to do n%2 == 1
		if n&1 == 1 {
			res *= x
		}
		x *= x    // Square the base
		n >>= 1   // Right shift: faster way to do n = n / 2
	}
	return res
}


func main() {
	fmt.Println(power(3, 13))
}