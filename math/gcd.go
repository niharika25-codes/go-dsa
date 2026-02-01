package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}


func gcd2(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a == b {
		return a
	}
	if a > b {
		if a%b == 0 {
			return b
		}
		return gcd(a%b, b)
	}
	if b%a == 0 {
		return a
	}
	return gcd(a, b%a)
}

func gcd_recur(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd_recur(b, a%b)
}

func main() {
	fmt.Println(gcd(98, 56))
}