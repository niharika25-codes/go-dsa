package main

import "fmt"

func SieveOfEratosthenes(n int) []int {
	falses := make([]bool, n+1)
	primes := []int{}

	for p := 2; p*p <= n; p++ {
		if !falses[p] {
			for i := p * p; i <= n; i += p {
				falses[i] = true
			}
		}
	}
	for p := 2; p <= n; p++ {
		if !falses[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	n := 30
	fmt.Printf("Primes up to %d: %v\n", n, SieveOfEratosthenes(n))
	n = 50
	fmt.Printf("Primes up to %d: %v\n", n, SieveOfEratosthenes(n))
}