package main

import "fmt"

func find_trailing_zeros(n int) int {
	if n < 0 {
		return -1
	}
	count := 0
	for i := 5; n/i >= 1; i*= 5 {
		count += n / i
	}
	return count
}

func main() {
	fmt.Println("zereos in 100: ", find_trailing_zeros(100))
}