package main

import "fmt"

func iterFact(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	fmt.Println(iterFact(12))
}