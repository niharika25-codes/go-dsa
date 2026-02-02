package main

import (
	"fmt"
	"math"
	"sort"
)

func getDivisors(n int) []int {
	var divisors []int
	sqrtN := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			// Avoid adding the square root twice
			if i*i != n {
				divisors = append(divisors, n/i)
			}
		}
	}
	sort.Ints(divisors)
	return divisors
}

func main() {
	number := 36
	fmt.Printf("Divisors of %d: %v\n", number, getDivisors(number))
}