package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func prime_factors(n int) {
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			for n%i == 0 {
				fmt.Print(i, " ")
				n /= i
			}
		}
	}
}

func main() {
	fmt.Println("Enter a number")
	a := 0
	fmt.Scan(&a)
	prime_factors(a)
}