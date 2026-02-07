package main

import (
	"fmt"
	"math"
	"time"
)

// Method 1: Optimized Trial Division
// Time Complexity: O(n^0.75) | Space: O(1)
func isPrimeTrial(num int) bool {
	if num < 2 { return false }
	if num == 2 || num == 3 { return true }
	if num%2 == 0 || num%3 == 0 { return false }
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func countWithTrialDivision(n int) int {
	limit := int(math.Sqrt(float64(n)))
	count := 0
	for i := 2; i <= limit; i++ {
		if isPrimeTrial(i) {
			count++
		}
	}
	return count
}

// Method 2: Sieve of Eratosthenes
// Time Complexity: O(sqrt(n) log log sqrt(n)) | Space: O(sqrt(n))
func countWithSieve(n int) int {
	limit := int(math.Sqrt(float64(n)))
	if limit < 2 { return 0 }

	primes := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= limit; p++ {
		if primes[p] {
			for i := p * p; i <= limit; i += p {
				primes[i] = false
			}
		}
	}

	count := 0
	for i := 2; i <= limit; i++ {
		if primes[i] {
			count++
		}
	}
	return count
}

func main() {
	n := 100000 // 10^5

	// Benchmark Trial Division
	startTrial := time.Now()
	resTrial := countWithTrialDivision(n)
	durationTrial := time.Since(startTrial)

	// Benchmark Sieve
	startSieve := time.Now()
	resSieve := countWithSieve(n)
	durationSieve := time.Since(startSieve)

	fmt.Printf("Results for n = %d\n", n)
	fmt.Println("---------------------------")
	fmt.Printf("Trial Division: %d numbers found in %v\n", resTrial, durationTrial)
	fmt.Printf("Sieve Method:   %d numbers found in %v\n", resSieve, durationSieve)
}